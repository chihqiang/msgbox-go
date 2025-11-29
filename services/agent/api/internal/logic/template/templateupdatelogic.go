// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package template

import (
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"

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
		return fmt.Errorf("not find agent")
	}
	var template models.Template
	if err := l.svcCtx.DB.Where("id = ? AND agent_id = ?", req.ID, agentID).First(&template).Error; err != nil {
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
	if req.Status != nil {
		template.Status = *req.Status
	}
	if req.Name != nil {
		template.Name = *req.Name
	}
	// 更新 Channels 关联
	if req.ChannelID != nil {
		var channel models.Channel
		if err := l.svcCtx.DB.Where("id = ?", *req.ChannelID).First(&channel).Error; err != nil {
			if gorm.ErrRecordNotFound == err {
				return errors.New("指定的通道不存在")
			}
			return err
		}
		template.ChannelID = *req.ChannelID
	}

	if err := l.svcCtx.DB.Model(&template).Where(models.Template{ID: req.ID, AgentID: agentID}).Updates(template).Error; err != nil {
		return err
	}
	return nil
}
