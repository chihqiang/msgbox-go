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

type ChannelUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelUpdateLogic {
	return &ChannelUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelUpdateLogic) ChannelUpdate(req *types.ChannelUpdateReq) error {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("not find agent")
	}
	updateData := models.Channel{}
	if req.Name != nil {
		updateData.Name = *req.Name
	}
	if req.VendorName != nil {
		updateData.VendorName = *req.VendorName
	}
	if req.Config != nil {
		config, err := json.Marshal(req.Config)
		if err != nil {
			return err
		}
		updateData.Config = config
	}
	if req.Status != nil {
		updateData.Status = *req.Status
	}
	if err := l.svcCtx.DB.Model(&models.Channel{}).Where("id = ? AND agent_id = ?", req.ID, agentID).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}
