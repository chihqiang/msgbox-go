package tasks

import (
	"chihqiang/msgbox-go/pkg/workflow"
	"chihqiang/msgbox-go/services/common/errs"
	"chihqiang/msgbox-go/services/common/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CheckTemplateTask struct {
	Log          logx.Logger
	DB           *gorm.DB
	TemplateCode string
}

func NewCheckTemplateTask(log logx.Logger, db *gorm.DB, templateCode string) *CheckTemplateTask {
	return &CheckTemplateTask{Log: log, DB: db, TemplateCode: templateCode}
}

func (c *CheckTemplateTask) Task() *workflow.Task {
	return &workflow.Task{
		Action: func(ctx context.Context) (context.Context, error) {
			if c.TemplateCode == "" {
				c.Log.Error("template code is empty")
				return ctx, errs.ErrTemplateCodeMissing
			}
			var template models.Template
			_ = c.DB.Model(template).Preload("Channel").Where(models.Template{Code: c.TemplateCode}).First(&template).Error
			if template.ID == 0 {
				c.Log.Error("template not found, template code: %s", c.TemplateCode)
				return ctx, errs.ErrTemplateCodeMissing
			}
			if template.Channel == nil {
				c.Log.Error("template channel not found, template code: %s", c.TemplateCode)
				return ctx, errs.ErrTemplateChannelMissing
			}
			ctx = context.WithValue(ctx, CtxModelTemplate, &template)
			ctx = context.WithValue(ctx, CtxModelChannel, template.Channel)
			return ctx, nil
		},
	}
}
