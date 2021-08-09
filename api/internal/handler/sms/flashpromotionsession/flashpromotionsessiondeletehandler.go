package handler

import (
	"net/http"

	"micro/api/internal/logic/sms/flashpromotionsession"
	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func FlashPromotionSessionDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteFlashPromotionSessionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFlashPromotionSessionDeleteLogic(r.Context(), ctx)
		resp, err := l.FlashPromotionSessionDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
