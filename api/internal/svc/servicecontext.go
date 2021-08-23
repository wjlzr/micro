package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"micro/api/internal/config"
	"micro/api/internal/middleware"
	"micro/rpc/svs/oauth/oauthclient"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	Oauth    oauthclient.Oauth
	Redis    *redis.Redis
	//Span     opentracing.Span
	//Tracer   opentracing.Tracer
}

func NewServiceContext(c config.Config) *ServiceContext {

	newRedis := redis.NewRedis(c.Redis.Address, redis.NodeType)

	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware(newRedis).Handle,
		Oauth:    oauthclient.NewOauth(zrpc.MustNewClient(c.OauthRpc)),
		Redis:    newRedis,
		//Span:     span,
		//Tracer:   tracer,
	}
}
