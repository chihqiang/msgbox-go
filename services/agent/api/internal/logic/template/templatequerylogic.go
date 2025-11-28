// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package template

import (
	"chihqiang/msgbox-go/pkg/timex"
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type TemplateQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTemplateQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateQueryLogic {
	return &TemplateQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TemplateQueryLogic) TemplateQuery(req *types.TemplateQueryReq) (resp *types.TemplateQueryResp, err error) {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return nil, fmt.Errorf("not find agent")
	}
	total, templates, err := models.NewPagination[models.Template](l.svcCtx.DB).QueryPage(req.Page, req.Size, func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("agent_id = ?", agentID)
		if req.Keywords != "" {
			keyword := "%" + req.Keywords + "%"
			tx = tx.Where("code LIKE ?", keyword).Or("vendor_code LIKE ?", keyword).Or("content LIKE ?", keyword)
		}
		return tx
	})
	if err != nil {
		return nil, err
	}
	return &types.TemplateQueryResp{
		Total: total,
		Data:  l.convert(templates),
	}, nil
}

func (l TemplateQueryLogic) convert(templates []models.Template) []types.TemplateItemResp {
	items := make([]types.TemplateItemResp, 0, len(templates))
	for _, item := range templates {
		items = append(items, types.TemplateItemResp{
			ID:         item.ID,
			AgentID:    item.AgentID,
			Name:       item.Name,
			Code:       item.Code,
			VendorCode: item.VendorCode,
			Signature:  item.Signature,
			Content:    item.Content,
			Status:     item.Status,
			UsedCount:  item.UsedCount,
			CreatedAt:  timex.FormatDate(item.CreatedAt),
			UpdatedAt:  timex.FormatDate(item.UpdatedAt),
			ChannelID:  item.ChannelID,
		})
	}
	return items
}
