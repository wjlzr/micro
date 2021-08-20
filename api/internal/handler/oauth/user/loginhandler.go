package user

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/tal-tech/go-zero/rest/httpx"
	logic "micro/api/internal/logic/oauth/user"
	"micro/api/internal/svc"
	"micro/api/internal/types"
	"net/http"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var header http.Header
		// 继承上下文关系，创建子 span
		childSpan := ctx.Tracer.StartSpan(
			"B",
			opentracing.ChildOf(ctx.Span.Context()),
		)
		ext.SpanKindRPCClient.Set(childSpan)
		_ = ctx.Tracer.Inject(childSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
		defer childSpan.Finish()

		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req, r.RemoteAddr)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
