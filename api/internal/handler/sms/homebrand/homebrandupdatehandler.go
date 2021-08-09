package handler

import (
	"net/http"

	"micro/api/internal/logic/sms/homebrand"
	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func HomeBrandUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateHomeBrandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHomeBrandUpdateLogic(r.Context(), ctx)
		resp, err := l.HomeBrandUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
