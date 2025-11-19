// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package template

import (
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type TemplateUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTemplateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateUpdateLogic {
	return &TemplateUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TemplateUpdateLogic) TemplateUpdate(req *types.TemplateUpdateReq) error {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("not find api")
	}
	var template models.Template
	if err := l.svcCtx.DB.Preload("Channels").Where("id = ? AND agent_id = ?", req.ID, agentID).First(&template).Error; err != nil {
		return err
	}
	if req.VendorCode != nil {
		template.VendorCode = *req.VendorCode
	}
	if req.Signature != nil {
		template.Signature = *req.Signature
	}
	if req.Content != nil {
		template.Content = *req.Content
	}
	if req.ContentType != nil {
		template.ContentType = *req.ContentType
	}
	if req.Status != nil {
		template.Status = *req.Status
	}
	// 更新 Channels 关联
	if req.ChannelID != nil {
		var channels []*models.Channel
		if len(req.ChannelID) > 0 {
			if err := l.svcCtx.DB.Where("id IN ?", req.ChannelID).Find(&channels).Error; err != nil {
				return err
			}
		}
		// 替换关联
		if err := l.svcCtx.DB.Association("Channels").Replace(channels); err != nil {
			return err
		}
	}
	return nil
}
