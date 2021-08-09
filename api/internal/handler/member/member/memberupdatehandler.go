package handler

import (
	"net/http"

	"micro/api/internal/logic/member/member"
	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MemberUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMemberUpdateLogic(r.Context(), ctx)
		resp, err := l.MemberUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
