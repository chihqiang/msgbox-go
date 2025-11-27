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
		return fmt.Errorf("not find agent")
	}
	var channel models.Channel
	if err := l.svcCtx.DB.Where("id = ? AND agent_id = ?", req.ChannelID, agentID).First(&channel).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			// 如果通道不存在，返回错误
			return errors.New("指定的通道不存在")
		}
		return err
	}
	template := &models.Template{
		AgentID:    agentID,
		ChannelID:  channel.ID,
		Code:       req.Code,
		VendorCode: req.VendorCode,
		Signature:  req.Signature,
		Content:    req.Content,
		Status:     req.Status,
	}
	if err := l.svcCtx.DB.Create(template).Error; err != nil {
		return err
	}
	return nil
}
