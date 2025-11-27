// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package agetent

import (
	"chihqiang/msgbox-go/pkg/timex"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"encoding/json"
	"fmt"

	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info() (resp *types.InfoResp, err error) {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return nil, fmt.Errorf("not find agent")
	}
	var agent models.Agent
	if err := l.svcCtx.DB.Model(&agent).First(&agent, agentID).Error; err != nil {
		return nil, err
	}
	return &types.InfoResp{
		ID:          agent.ID,
		AgentNo:     agent.AgentNo,
		AgentSecret: agent.AgentSecret,
		Name:        agent.Name,
		Phone:       agent.Phone,
		Email:       agent.Email,
		Status:      agent.Status,
		CreatedAt:   timex.FormatDate(agent.CreatedAt),
		UpdatedAt:   timex.FormatDate(agent.UpdatedAt),
	}, nil
}
