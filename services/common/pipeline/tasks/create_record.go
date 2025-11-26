package tasks

import (
	"chihqiang/msgbox-go/pkg/stringx"
	"chihqiang/msgbox-go/pkg/workflow"
	"chihqiang/msgbox-go/services/common/errs"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateRecordTask struct {
	Log       logx.Logger
	DB        *gorm.DB
	Receivers []string
	Variables map[string]string
	Extra     map[string]interface{}
}

func NewCreateRecordTask(log logx.Logger, db *gorm.DB, receivers []string, variables map[string]string, extra map[string]interface{}) *CreateRecordTask {
	crt := &CreateRecordTask{Log: log, DB: db, Receivers: receivers, Variables: variables, Extra: extra}
	return crt
}

func (c *CreateRecordTask) Task() *workflow.Task {
	return &workflow.Task{
		Action: func(ctx context.Context) (context.Context, error) {
			now := time.Now()
			batch := models.SendBatch{
				BatchNo:       stringx.UUID(),
				TotalCount:    len(c.Receivers),
				Status:        models.SendBatchStatusPending,
				ScheduledTime: &now,
				Agent:         ctx.Value(CtxModelAgent).(*models.Agent),
				Channel:       ctx.Value(CtxModelChannel).(*models.Channel),
				Template:      ctx.Value(CtxModelTemplate).(*models.Template),
			}
			for _, s := range c.Receivers {
				content := strings.Join([]string{
					batch.Template.Signature,
					stringx.ReplaceVariables(batch.Template.Content, c.Variables),
				}, "")
				rc := &models.SendRecord{
					Receiver:      s,
					VendorName:    batch.Channel.VendorName,
					ChannelConfig: batch.Channel.Config,
					VendorCode:    batch.Template.VendorCode,
					Signature:     batch.Template.Signature,
					Content:       content,
					Variables:     models.MapToDataTypesJSON(c.Variables),
					Extra:         models.MapToDataTypesJSON(c.Extra),
					Status:        models.SendRecordStatusPending,
					Agent:         batch.Agent,
					Channel:       batch.Channel,
					Template:      batch.Template,
					Batch:         &batch,
				}
				batch.Records = append(batch.Records, rc)
			}

			err := c.DB.Create(&batch).Error
			if err != nil {
				c.Log.Error("create send batch failed, err: %v", err)
				return ctx, errs.ErrDB
			}
			ctx = context.WithValue(ctx, CtxModelSendBatch, &batch)
			ctx = context.WithValue(ctx, CtxModelSendRecord, batch.Records)
			return ctx, nil
		},
	}
}
