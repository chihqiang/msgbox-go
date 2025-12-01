// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package agetent

import (
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"github.com/samber/lo"

	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetSecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetSecretLogic {
	return &ResetSecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetSecretLogic) ResetSecret() (resp *types.ResetSecretResp, err error) {
	agentID, err := types.GetAgentID(l.ctx)
	if err != nil {
		return nil, err
	}
	agentSecret := lo.RandomString(32, lo.AlphanumericCharset)
	if err := l.svcCtx.DB.Model(&models.Agent{}).Where(models.Agent{ID: agentID}).Updates(&models.Agent{AgentSecret: agentSecret}).Error; err != nil {
		return nil, err
	}
	return &types.ResetSecretResp{AgentSecret: agentSecret}, nil
}
