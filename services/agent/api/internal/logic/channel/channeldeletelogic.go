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

type ChannelDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelDeleteLogic {
	return &ChannelDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelDeleteLogic) ChannelDelete(req *types.IDReq) error {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("not find api")
	}
	var channel models.Channel
	count := l.svcCtx.DB.Model(&channel).Association("Templates").Count()
	if count > 0 {
		return fmt.Errorf("channel is used by %d template(s), cannot delete", count)
	}
	if err := l.svcCtx.DB.Model(&models.Channel{}).Where("id = ? AND agent_id = ?", req.ID, agentID).Delete(&models.Channel{}).Error; err != nil {
		return err
	}
	return nil
}
