package handler

import (
	"net/http"

	"github.com/jacksonyoudi/datacenter/internal/logic/user"
	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func PingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewPingLogic(r.Context(), ctx)
		err := l.Ping()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
