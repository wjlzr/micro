package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"micro/rpc/model/paymodel"
	"micro/rpc/svs/pay/internal/config"
)

type ServiceContext struct {
	c                config.Config
	WxRecordModel    paymodel.PayWxRecordModel
	WxMerchantsModel paymodel.PayWxMerchantsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c:                c,
		WxRecordModel:    paymodel.NewPayWxRecordModel(sqlConn),
		WxMerchantsModel: paymodel.NewPayWxMerchantsModel(sqlConn),
	}
}
