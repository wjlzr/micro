package handler

import (
	"net/http"

	"micro/api/internal/logic/sms/flashpromotionsession"
	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func FlashPromotionSessionAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddFlashPromotionSessionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFlashPromotionSessionAddLogic(r.Context(), ctx)
		resp, err := l.FlashPromotionSessionAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
