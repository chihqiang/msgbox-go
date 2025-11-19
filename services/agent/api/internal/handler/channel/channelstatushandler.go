// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package channel

import (
	"chihqiang/msgbox-go/services/agent/api/internal/logic/channel"
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func ChannelStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := channel.NewChannelStatusLogic(r.Context(), svcCtx)
		err := l.ChannelStatus(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
