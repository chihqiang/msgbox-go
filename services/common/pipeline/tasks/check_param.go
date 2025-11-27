package tasks

import (
	"chihqiang/msgbox-go/pkg/workflow"
	"chihqiang/msgbox-go/services/common/errs"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckParamTask struct {
	Log          logx.Logger
	AgentNo      string
	AgentSecret  string
	TemplateCode string
	Receivers    []string
	Variables    map[string]string
}

func NewCheckParamTask(log logx.Logger, agentNo string, agentSecret string, templateCode string, receivers []string, variables map[string]string) *CheckParamTask {
	cpt := &CheckParamTask{
		Log:          log,
		AgentNo:      agentNo,
		AgentSecret:  agentSecret,
		TemplateCode: templateCode,
		Receivers:    receivers,
		Variables:    variables,
	}
	return cpt
}

func (c *CheckParamTask) Task() *workflow.Task {
	return &workflow.Task{
		Action: func(ctx context.Context) (context.Context, error) {
			if c.AgentNo == "" {
				c.Log.Error("agent no is empty")
				return ctx, errs.ErrParamInvalid
			}
			if c.AgentSecret == "" {
				c.Log.Error("agent key is empty")
				return ctx, errs.ErrParamInvalid
			}
			if c.TemplateCode == "" {
				c.Log.Error("template code is empty")
				return ctx, errs.ErrParamInvalid
			}
			if len(c.Receivers) == 0 {
				c.Log.Error("receivers is empty")
				return ctx, errs.ErrParamInvalid
			}
			return ctx, nil
		},
	}
}
