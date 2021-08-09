package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"micro/rpc/model/omsmodel"
	"micro/rpc/svs/oms/internal/config"
)

type ServiceContext struct {
	c config.Config

	OmsCartItemModel            *omsmodel.OmsCartItemModel
	OmsCompanyAddressModel      *omsmodel.OmsCompanyAddressModel
	OmsOrderItemModel           *omsmodel.OmsOrderItemModel
	OmsOrderModel               *omsmodel.OmsOrderModel
	OmsOrderOperateHistoryModel *omsmodel.OmsOrderOperateHistoryModel
	OmsOrderReturnApplyModel    *omsmodel.OmsOrderReturnApplyModel
	OmsOrderReturnReasonModel   *omsmodel.OmsOrderReturnReasonModel
	OmsOrderSettingModel        *omsmodel.OmsOrderSettingModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c: c,

		OmsCartItemModel:            omsmodel.NewOmsCartItemModel(sqlConn),
		OmsCompanyAddressModel:      omsmodel.NewOmsCompanyAddressModel(sqlConn),
		OmsOrderItemModel:           omsmodel.NewOmsOrderItemModel(sqlConn),
		OmsOrderModel:               omsmodel.NewOmsOrderModel(sqlConn),
		OmsOrderOperateHistoryModel: omsmodel.NewOmsOrderOperateHistoryModel(sqlConn),
		OmsOrderReturnApplyModel:    omsmodel.NewOmsOrderReturnApplyModel(sqlConn),
		OmsOrderReturnReasonModel:   omsmodel.NewOmsOrderReturnReasonModel(sqlConn),
		OmsOrderSettingModel:        omsmodel.NewOmsOrderSettingModel(sqlConn),
	}
}
