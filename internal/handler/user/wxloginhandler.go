package handler

import (
	"net/http"

	"github.com/jacksonyoudi/datacenter/internal/logic/user"
	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func WxloginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WxLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWxloginLogic(r.Context(), ctx)
		resp, err := l.Wxlogin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
