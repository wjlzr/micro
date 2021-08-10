package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"micro/rpc/model/oauth"
	"micro/rpc/svs/oauth/internal/config"
	"os"
	"time"
)

type ServiceContext struct {
	Config   config.Config
	Redis    *redis.Redis
	Eloquent *gorm.DB
	Oauth    *oauth.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	newRedis := redis.NewRedis(c.Redis.Address, redis.NodeType)

	// mysql-orm集成gorm
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	Eloquent, _ := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{Logger: newLogger})

	return &ServiceContext{
		Config:   c,
		Redis:    newRedis,
		Eloquent: Eloquent,
		Oauth:    oauth.NewSysUserModel(Eloquent),
	}
}
