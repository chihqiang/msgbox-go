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

type TemplateCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTemplateCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateCreateLogic {
	return &TemplateCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TemplateCreateLogic) TemplateCreate(req *types.TemplateCreateReq) error {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("not find api")
	}
	template := &models.Template{
		AgentID:     agentID,
		Code:        req.Code,
		VendorCode:  req.VendorCode,
		Signature:   req.Signature,
		Content:     req.Content,
		ContentType: req.ContentType,
		Status:      req.Status,
	}
	var channels []*models.Channel
	if err := l.svcCtx.DB.Where("id IN ?", req.ChannelID).Find(&channels).Error; err != nil {
		return err
	}
	template.Channels = channels
	if err := l.svcCtx.DB.Create(template).Error; err != nil {
		return err
	}
	return nil
}
