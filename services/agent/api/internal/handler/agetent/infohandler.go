// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package agetent

import (
	"net/http"

	"chihqiang/msgbox-go/services/agent/api/internal/logic/agetent"
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	xhttp "github.com/zeromicro/x/http"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := agetent.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
