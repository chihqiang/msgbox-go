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

type ChannelCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelCreateLogic {
	return &ChannelCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelCreateLogic) ChannelCreate(req *types.ChannelCreateReq) error {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("not find api")
	}
	var count int64
	_ = l.svcCtx.DB.Model(&models.Channel{}).Where(&models.Channel{
		AgentID: agentID,
		Code:    req.Code,
	}).Count(&count).Error
	if count > 0 {
		return fmt.Errorf("%s模版已存在", req.Code)
	}
	if err = l.svcCtx.DB.Model(&models.Channel{}).Create(
		&models.Channel{
			AgentID:    agentID,
			Code:       req.Code,
			Name:       req.Name,
			VendorName: req.VendorName,
			Config:     models.MapToDataTypesJSON(req.Config),
			Status:     req.Status,
		},
	).Error; err != nil {
		return err
	}
	return nil
}
