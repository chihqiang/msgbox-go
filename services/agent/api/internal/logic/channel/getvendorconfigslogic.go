// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package channel

import (
	"chihqiang/msgbox-go/pkg/htmlx"
	"chihqiang/msgbox-go/services/common/channels"
	"context"

	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVendorConfigsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVendorConfigsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVendorConfigsLogic {
	return &GetVendorConfigsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVendorConfigsLogic) GetVendorConfigs() (resp []types.GetVendorConfigsResp, err error) {
	senders := channels.GetSenders()
	for _, sender := range senders {
		formFields := make([]types.FormField, 0)
		for _, form := range htmlx.ToFormFields(sender) {
			formFields = append(formFields, types.FormField{
				Type:        form.Type,
				Name:        form.Name,
				Label:       form.Label,
				Required:    form.Required,
				Placeholder: form.Placeholder,
				Default:     form.Default,
			})
		}
		resp = append(resp, types.GetVendorConfigsResp{
			Name:    sender.GetName(),
			Label:   sender.GetLabel(),
			Configs: formFields,
		})
	}
	return
}
