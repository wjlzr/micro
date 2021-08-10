package svc

import (
	"github.com/k0kubun/pp"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"micro/rpc/model/omsmodel"
	"micro/rpc/svs/oms/internal/config"
	"os"
	"time"
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

	// mysql-orm集成gorm
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	Eloquent, _ := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{Logger: newLogger})
	pp.Println("A服务初始化mysql服务")
	pp.Println(Eloquent)
	return &ServiceContext{
		c:                           c,
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
