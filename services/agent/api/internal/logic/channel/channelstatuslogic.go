// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package channel

import (
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelStatusLogic {
	return &ChannelStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelStatusLogic) ChannelStatus(req *types.IDStatusReq) error {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("not find agent")
	}
	if err := l.svcCtx.DB.Model(&models.Channel{}).Where("id = ? AND agent_id = ?", req.ID, agentID).Update("status", req.Status).Error; err != nil {
		return err
	}
	return nil
}
