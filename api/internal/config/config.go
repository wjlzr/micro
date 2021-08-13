package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// oauth
	OauthRpc zrpc.RpcClientConf

	// Auth
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	// Redis
	Redis struct {
		Address string
	}
}
