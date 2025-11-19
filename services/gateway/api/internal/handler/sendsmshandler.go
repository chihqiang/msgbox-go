// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"chihqiang/msgbox-go/services/gateway/api/internal/logic"
	"chihqiang/msgbox-go/services/gateway/api/internal/svc"
	"chihqiang/msgbox-go/services/gateway/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func SendSmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendRequest
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendSmsLogic(r.Context(), svcCtx)
		resp, err := l.SendSms(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
