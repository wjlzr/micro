package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"micro/api/internal/config"
	"micro/api/internal/middleware"
	"micro/rpc/svs/oms/omsclient"
	"micro/rpc/svs/pms/pmsclient"
	"micro/rpc/svs/sms/smsclient"
	"micro/rpc/svs/sys/sysclient"
	"micro/rpc/svs/ums/umsclient"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	Sys      sysclient.Sys
	Ums      umsclient.Ums
	Pms      pmsclient.Pms
	Oms      omsclient.Oms
	Sms      smsclient.Sms
	Redis    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.NewRedis(c.Redis.Address, redis.NodeType)
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware(newRedis).Handle,
		Sys:      sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
		Ums:      umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		Pms:      pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		Oms:      omsclient.NewOms(zrpc.MustNewClient(c.OmsRpc)),
		Sms:      smsclient.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		Redis:    newRedis,
	}
}
