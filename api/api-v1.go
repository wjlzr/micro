package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tal-tech/go-zero/rest/httpx"
	"micro/api/internal/config"
	"micro/api/internal/handler"
	"micro/api/internal/svc"
	"micro/common/response"
	"net/http"
	"time"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var (
	configFile   = flag.String("f", "api/etc/api-v1.yaml", "the config file")
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	flag.Parse()
	recordMetrics()
	// 链路追踪
	//tracer, closer, _ := trace.CreateTracer("api-v1-service")
	//parentSpan := tracer.StartSpan("A")

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 结束 A
	//parentSpan.Finish()
	//// 结束当前 tracer
	//_ = closer.Close()

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *response.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
