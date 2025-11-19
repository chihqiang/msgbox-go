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

type TemplateDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTemplateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateDeleteLogic {
	return &TemplateDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TemplateDeleteLogic) TemplateDelete(req *types.IDReq) error {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("not find api")
	}
	var template models.Template
	// 先查询模板（包括关联 Channels）
	if err := l.svcCtx.DB.Preload("Channels").Where("id = ? AND agent_id = ?", req.ID, agentID).First(&template).Error; err != nil {
		return err
	}
	// 删除多对多关联
	if err := l.svcCtx.DB.Model(&template).Association("Channels").Clear(); err != nil {
		return err
	}
	// 删除模板记录（软删除，如果需要硬删除用 Unscoped().Delete）
	if err := l.svcCtx.DB.Delete(&template).Error; err != nil {
		return err
	}
	return nil
}
