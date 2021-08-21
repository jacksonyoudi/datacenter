package handler

import (
	"net/http"

	"github.com/jacksonyoudi/datacenter/internal/logic/user"
	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func Code2SessionHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewCode2SessionLogic(r.Context(), ctx)
		resp, err := l.Code2Session()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
