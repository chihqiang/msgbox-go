package pipeline

import (
	"chihqiang/msgbox-go/pkg/workflow"
	"chihqiang/msgbox-go/services/common/errs"
	"chihqiang/msgbox-go/services/common/models"
	"chihqiang/msgbox-go/services/common/pipeline/tasks"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ITask interface {
	Task() workflow.TaskInterface
}
type SendPipeline struct {
	Log          logx.Logger
	DB           *gorm.DB
	TraceID      string
	AgentNo      string
	AgentSecret  string
	TemplateCode string
	Receivers    []string
	Variables    map[string]string
	Extra        map[string]interface{}
	sendBatch    *models.SendBatch
}

func (p *SendPipeline) Check(ctx context.Context) error {
	serial := workflow.NewStageSerial()
	serial.Add(tasks.NewCheckParamTask(p.Log, p.AgentNo, p.AgentSecret, p.TemplateCode, p.Receivers, p.Variables).Task())
	serial.Add(tasks.NewCheckAgentTask(p.Log, p.DB, p.AgentNo, p.AgentSecret).Task())
	serial.Add(tasks.NewCheckTemplateTask(p.Log, p.DB, p.TemplateCode).Task())
	serial.Add(tasks.NewCreateRecordTask(p.Log, p.DB, p.TraceID, p.Receivers, p.Variables, p.Extra).Task())
	serial.Add(&workflow.Task{
		Action: func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		},
		Finish: func(ctx context.Context, err error) {
			if err != nil {
				p.Log.Error("check send pipeline failed, err: %v", err)
				return
			}
			if batch, ok := ctx.Value(tasks.CtxModelSendBatch).(*models.SendBatch); ok {
				p.sendBatch = batch
			}
		},
	})
	return serial.Run(ctx)
}

func (p *SendPipeline) Send(ctx context.Context) error {
	if p.sendBatch == nil {
		p.Log.Error("send batch is nil, check must be run first")
		return fmt.Errorf("send batch is nil, check must be run first")
	}
	serial := workflow.NewStageSerial()
	// 开始发送
	serial.Add(&workflow.Task{
		Action: func(ctx context.Context) (context.Context, error) {
			now := time.Now()
			if err := p.DB.Model(&models.SendBatch{}).
				Where(&models.SendBatch{ID: p.sendBatch.ID}).
				Updates(&models.SendBatch{SendStartTime: &now}).Error; err != nil {
				p.Log.Error("update send batch start time failed, err: %v", err)
				return ctx, errs.ErrDB
			}
			return ctx, nil
		},
	})
	//发送中
	serial.Add(&workflow.Task{
		Action: func(ctx context.Context) (context.Context, error) {
			parallel := workflow.NewStageParallel()
			for _, record := range p.sendBatch.Records {
				parallel.Add(tasks.NewSendTask(p.Log, p.DB, record).Task())
			}
			_ = parallel.Run(ctx)
			return ctx, nil
		},
	})
	//结束发送
	serial.Add(&workflow.Task{
		Action: func(ctx context.Context) (context.Context, error) {
			now := time.Now()
			if err := p.DB.Model(&models.SendBatch{}).
				Where(&models.SendBatch{ID: p.sendBatch.ID}).
				Updates(&models.SendBatch{SendEndTime: &now}).Error; err != nil {
				p.Log.Error("update send batch end time failed, err: %v", err)
				return ctx, errs.ErrDB
			}
			return ctx, nil
		},
	})
	return serial.Run(ctx)
}
func (p *SendPipeline) GetSendBatch() (*models.SendBatch, error) {
	if p.sendBatch == nil {
		p.Log.Error("send batch is nil, send must be run first")
		return nil, fmt.Errorf("send batch is nil")
	}
	var batch models.SendBatch
	err := p.DB.Preload("Records").First(&batch, p.sendBatch.ID).Error
	if err != nil {
		p.Log.Error("get send batch failed, err: %v", err)
		return nil, err
	}
	p.sendBatch = &batch
	return p.sendBatch, nil
}
