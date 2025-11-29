package tasks

import (
	"chihqiang/msgbox-go/pkg/workflow"
	"chihqiang/msgbox-go/services/common/channels/senders"
	"chihqiang/msgbox-go/services/common/errs"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type SendTask struct {
	Log    logx.Logger
	DB     *gorm.DB
	record *models.SendRecord
}

func NewSendTask(log logx.Logger, db *gorm.DB, record *models.SendRecord) *SendTask {
	return &SendTask{
		Log:    log,
		DB:     db,
		record: record,
	}
}

func (s *SendTask) getSender() (senders.ISender, error) {
	for _, sender := range senders.Get() {
		if sender.GetName() == s.record.VendorName {
			if err := sender.SetConfig(models.DataTypesToMap(s.record.ChannelConfig)); err != nil {
				return nil, err
			}
			return sender, nil
		}
	}
	return nil, fmt.Errorf("sender not found for vendor: %s", s.record.VendorName)
}

func (s *SendTask) Task() *workflow.Task {
	return &workflow.Task{
		Action: func(ctx context.Context) (context.Context, error) {
			sender, err := s.getSender()
			if err != nil {
				s.Log.Error("get sender failed, err: %v", err)
				return ctx, s.fail(err, map[string]any{})
			}
			resp, err := sender.Send(s.record)
			if err != nil {
				s.Log.Error("send message failed, err: %v", err)
				return ctx, s.fail(err, resp)
			}
			return ctx, s.success(resp)
		},
	}
}

func (s *SendTask) success(response map[string]any) error {
	now := time.Now()
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// 更新发送记录
		if err := tx.Model(&s.record).Updates(map[string]interface{}{
			"send_time": now,
			"status":    models.SendRecordStatusSending,
			"response":  models.MapToDataTypesJSON(response),
		}).Error; err != nil {
			s.Log.Error("update send record failed, err: %v", err)
			return err
		}
		// 更新批次统计
		if err := tx.Model(&models.SendBatch{}).
			Where("id = ?", s.record.BatchID).
			UpdateColumn("success_count", gorm.Expr("success_count + ?", 1)).Error; err != nil {
			s.Log.Error("update send batch success count failed, err: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		return errs.ErrDB
	}
	return nil
}

func (s *SendTask) fail(errMsg error, response map[string]any) error {
	now := time.Now()
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// 更新发送记录
		if err := tx.Model(&s.record).Updates(map[string]interface{}{
			"send_time": now,
			"status":    models.SendRecordStatusFailed,
			"error":     errMsg.Error(),
			"response":  models.MapToDataTypesJSON(response),
		}).Error; err != nil {
			s.Log.Error("update send record failed, err: %v", err)
			return errs.ErrDB
		}
		if err := tx.Model(&models.SendBatch{}).
			Where("id = ?", s.record.BatchID).
			UpdateColumn("fail_count", gorm.Expr("fail_count + ?", 1)).Error; err != nil {
			s.Log.Error("update send batch fail count failed, err: %v", err)
			return errs.ErrDB
		}
		return nil
	})
	if err != nil {
		return errs.ErrDB
	}
	return nil
}
