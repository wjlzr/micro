package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	JWT struct {
		AccessSecret string
		AccessExpire int64
	}

	Mysql struct {
		Dns string
	}
}
