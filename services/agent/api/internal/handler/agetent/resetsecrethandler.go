// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package agetent

import (
	"net/http"

	"chihqiang/msgbox-go/services/agent/api/internal/logic/agetent"
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	xhttp "github.com/zeromicro/x/http"
)

func ResetSecretHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := agetent.NewResetSecretLogic(r.Context(), svcCtx)
		resp, err := l.ResetSecret()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
