package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"micro/api/internal/config"
	"micro/api/internal/middleware"
	"micro/rpc/svs/oms/omsclient"
	"micro/rpc/svs/pms/pmsclient"
	"micro/rpc/svs/sms/smsclient"
	"micro/rpc/svs/sys/sysclient"
	"micro/rpc/svs/ums/umsclient"
	"os"
	"time"
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
	Eloquent *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.NewRedis(c.Redis.Address, redis.NodeType)

	// mysql-orm集成gorm
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	Eloquent, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{Logger: newLogger})
	if err != nil {

	}

	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware(newRedis).Handle,
		Sys:      sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
		Ums:      umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		Pms:      pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		Oms:      omsclient.NewOms(zrpc.MustNewClient(c.OmsRpc)),
		Sms:      smsclient.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		Redis:    newRedis,
		Eloquent: Eloquent,
	}
}
