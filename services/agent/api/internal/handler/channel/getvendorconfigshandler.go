// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package channel

import (
	"net/http"

	"chihqiang/msgbox-go/services/agent/api/internal/logic/channel"
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	xhttp "github.com/zeromicro/x/http"
)

func GetVendorConfigsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := channel.NewGetVendorConfigsLogic(r.Context(), svcCtx)
		resp, err := l.GetVendorConfigs()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
