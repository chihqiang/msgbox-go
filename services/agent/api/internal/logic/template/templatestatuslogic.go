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

type TemplateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTemplateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateStatusLogic {
	return &TemplateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TemplateStatusLogic) TemplateStatus(req *types.IDStatusReq) error {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("not find api")
	}
	if err := l.svcCtx.DB.Model(&models.Template{}).Where("id = ? AND agent_id = ?", req.ID, agentID).Update("status", req.Status).Error; err != nil {
		return err
	}
	return nil
}
