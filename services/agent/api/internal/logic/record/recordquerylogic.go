// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package record

import (
	"chihqiang/msgbox-go/pkg/timex"
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecordQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecordQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordQueryLogic {
	return &RecordQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecordQueryLogic) RecordQuery(req *types.RecordQueryReq) (resp *types.RecordQueryResp, err error) {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return nil, fmt.Errorf("not find agent")
	}

	db := l.svcCtx.DB.Model(&models.SendRecord{}).Preload("Channel").Where("agent_id = ?", agentID)
	if req.Keywords != "" {
		keyword := "%" + req.Keywords + "%"
		db = db.Where("receiver LIKE ?", keyword)
	}

	total, sendRecords, err := models.Page[models.SendRecord](db, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return &types.RecordQueryResp{
		Total: total,
		Data:  l.convert(sendRecords),
	}, nil
}

func (l RecordQueryLogic) convert(channels []models.SendRecord) []types.RecordItemResp {
	items := make([]types.RecordItemResp, 0, len(channels))
	for _, item := range channels {
		items = append(items, types.RecordItemResp{
			ID:            item.ID,
			Receiver:      item.Receiver,
			TraceID:       item.TraceID,
			ChannelName:   item.Channel.Name,
			ChannelConfig: models.DataTypesToMap(item.ChannelConfig),
			VendorName:    item.VendorName,
			VendorCode:    item.VendorCode,
			Signature:     item.Signature,
			Title:         item.Title,
			Content:       item.Content,
			Variables:     models.DataTypesToMap(item.Variables),
			Extra:         models.DataTypesToMap(item.Extra),
			Status:        item.Status,
			StatusMsg:     item.StatusMsg(),
			SendTime:      timex.FormatDate(item.SendTime),
			Error:         item.Error,
			Response:      models.DataTypesToMap(item.Response),
			DeliveryTime:  timex.FormatDate(item.DeliveryTime),
			DeliveryRaw:   models.DataTypesToMap(item.DeliveryRaw),
			CreatedAt:     timex.FormatDate(item.CreatedAt),
			UpdatedAt:     timex.FormatDate(item.UpdatedAt),
		})
	}
	return items
}
