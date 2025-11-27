// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package nologin

import (
	"net/http"

	"chihqiang/msgbox-go/services/agent/api/internal/logic/nologin"
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	xhttp "github.com/zeromicro/x/http"
)

func GetChannelConfigsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := nologin.NewGetChannelConfigsLogic(r.Context(), svcCtx)
		resp, err := l.GetChannelConfigs()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
