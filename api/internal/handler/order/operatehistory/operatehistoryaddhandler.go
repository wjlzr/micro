package handler

import (
	"net/http"

	"micro/api/internal/logic/order/operatehistory"
	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func OperateHistoryAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddOperateHistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewOperateHistoryAddLogic(r.Context(), ctx)
		resp, err := l.OperateHistoryAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
