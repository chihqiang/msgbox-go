package tasks

import (
	"chihqiang/msgbox-go/pkg/workflow"
	"chihqiang/msgbox-go/services/common/errs"
	"chihqiang/msgbox-go/services/common/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CheckAgentTask struct {
	Log      logx.Logger
	DB       *gorm.DB
	AgentNo  string
	AgentKey string
}

func NewCheckAgentTask(log logx.Logger, db *gorm.DB, agentNo string, agentKey string) *CheckAgentTask {
	return &CheckAgentTask{
		Log:      log,
		DB:       db,
		AgentNo:  agentNo,
		AgentKey: agentKey,
	}
}

func (c *CheckAgentTask) Task() *workflow.Task {
	return &workflow.Task{
		Action: func(ctx context.Context) (context.Context, error) {
			if c.AgentNo == "" || c.AgentKey == "" {
				c.Log.Error("agent no or agent key is empty")
				return ctx, errs.ErrAuthInvalid
			}
			var agent models.Agent
			_ = c.DB.Model(agent).Where(models.Agent{AgentNo: c.AgentNo, AgentKey: c.AgentKey}).First(&agent).Error
			if agent.ID == 0 {
				c.Log.Error("agent not found, agent no: %s, agent key: %s", c.AgentNo, c.AgentKey)
				return ctx, errs.ErrAuthInvalid
			}
			ctx = context.WithValue(ctx, CtxModelAgent, &agent)
			return ctx, nil
		},
	}
}
