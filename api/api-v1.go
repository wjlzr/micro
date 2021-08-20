package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"micro/common/errorx"
	"micro/services/trace"
	"net/http"

	"micro/api/internal/config"
	"micro/api/internal/handler"
	"micro/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "api/etc/api-v1.yaml", "the config file")

func main() {
	flag.Parse()

	// 链路追踪
	tracer, closer, _ := trace.CreateTracer("api-v1-service")
	parentSpan := tracer.StartSpan("A")

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c, tracer, parentSpan)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 结束 A
	parentSpan.Finish()
	// 结束当前 tracer
	_ = closer.Close()

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
