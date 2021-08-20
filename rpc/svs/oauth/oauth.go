package main

import (
	"flag"
	"fmt"
	"micro/rpc/svs/oauth/internal/config"
	"micro/rpc/svs/oauth/internal/server"
	"micro/rpc/svs/oauth/internal/svc"
	"micro/rpc/svs/oauth/oauth"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/oauth.yaml", "the config file")

func main() {
	flag.Parse()

	//// 链路追踪
	//var header http.Header
	//tracer, spanContext, closer, _ := trace.CreateSvsTracer("svs-oauth-service", header)
	//defer closer.Close()
	//// 生成依赖关系，并新建一个 span、
	//// 这里很重要，因为生成了  References []SpanReference 依赖关系
	//startSpan := tracer.StartSpan("operationName", ext.RPCServerOption(spanContext))
	//defer startSpan.Finish()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewOauthServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		oauth.RegisterOauthServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
