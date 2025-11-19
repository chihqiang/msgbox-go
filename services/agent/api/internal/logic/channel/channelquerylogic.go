// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package channel

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

type ChannelQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelQueryLogic {
	return &ChannelQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelQueryLogic) ChannelQuery(req *types.ChannelQueryReq) (resp *types.ChannelQueryResp, err error) {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return nil, fmt.Errorf("not find api")
	}

	total, channels, err := models.NewPagination[models.Channel](l.svcCtx.DB).QueryPage(req.Page, req.Size, func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("agent_id = ?", agentID)
		return tx
	})
	if err != nil {
		return nil, err
	}
	return &types.ChannelQueryResp{
		Total: total,
		Data:  l.convert(channels),
	}, nil
}

func (l ChannelQueryLogic) convert(channels []models.Channel) []types.ChannelItemResp {
	items := make([]types.ChannelItemResp, 0, len(channels))
	for _, item := range channels {
		var config map[string]string
		if err := json.Unmarshal(item.Config, &config); err != nil {
			continue
		}
		items = append(items, types.ChannelItemResp{
			ID:           item.ID,
			AgentID:      item.AgentID,
			Code:         item.Code,
			Name:         item.Name,
			ProviderCode: item.ProviderCode,
			Endpoint:     item.Endpoint,
			Config:       config,
			Status:       item.Status,
			CreatedAt:    timex.FormatDate(item.CreatedAt),
			UpdatedAt:    timex.FormatDate(item.UpdatedAt),
		})
	}
	return items
}
