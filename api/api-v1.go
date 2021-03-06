package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"micro/api/internal/config"
	"micro/api/internal/handler"
	"micro/api/internal/svc"
	"micro/common/response"
	"net/http"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var (
	configFile = flag.String("f", "api/etc/api-v1.yaml", "the config file")
)

func main() {

	// 日常打卡
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *response.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
