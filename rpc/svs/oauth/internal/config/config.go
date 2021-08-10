package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	JWT struct {
		AccessSecret string
		AccessExpire int64
	}

	Redis struct {
		Address string
	}

	Mysql struct {
		Dsn string
	}
}
