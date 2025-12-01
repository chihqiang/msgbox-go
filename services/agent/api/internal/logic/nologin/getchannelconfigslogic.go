// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package nologin

import (
	"chihqiang/msgbox-go/services/common/channels/senders"
	"context"

	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelConfigsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChannelConfigsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChannelConfigsLogic {
	return &GetChannelConfigsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChannelConfigsLogic) GetChannelConfigs() (resp []types.GetChannelConfigsResp, err error) {
	for _, sender := range senders.List() {
		formFields := make([]types.FormField, 0)
		for _, form := range sender.SenderFormFields() {
			formFields = append(formFields, types.FormField{
				Type:        form.Type,
				Name:        form.Name,
				Label:       form.Label,
				Required:    form.Required,
				Placeholder: form.Placeholder,
				Default:     form.Default,
			})
		}
		resp = append(resp, types.GetChannelConfigsResp{
			Name:    sender.Name,
			Label:   sender.Label,
			Configs: formFields,
		})
	}
	return
}
