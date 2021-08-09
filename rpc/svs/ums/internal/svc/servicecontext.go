package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"micro/rpc/model/umsmodel"
	"micro/rpc/svs/ums/internal/config"
)

type ServiceContext struct {
	c config.Config

	UmsGrowthChangeHistoryModel           *umsmodel.UmsGrowthChangeHistoryModel
	UmsIntegrationChangeHistoryModel      *umsmodel.UmsIntegrationChangeHistoryModel
	UmsIntegrationConsumeSettingModel     *umsmodel.UmsIntegrationConsumeSettingModel
	UmsMemberModel                        *umsmodel.UmsMemberModel
	UmsMemberLevelModel                   *umsmodel.UmsMemberLevelModel
	UmsMemberLoginLogModel                *umsmodel.UmsMemberLoginLogModel
	UmsMemberMemberTagRelationModel       *umsmodel.UmsMemberMemberTagRelationModel
	UmsMemberProductCategoryRelationModel *umsmodel.UmsMemberProductCategoryRelationModel
	UmsMemberReceiveAddressModel          *umsmodel.UmsMemberReceiveAddressModel
	UmsMemberRuleSettingModel             *umsmodel.UmsMemberRuleSettingModel
	UmsMemberStatisticsInfoModel          *umsmodel.UmsMemberStatisticsInfoModel
	UmsMemberTaskModel                    *umsmodel.UmsMemberTaskModel
	UmsMemberTagModel                     *umsmodel.UmsMemberTagModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c: c,

		UmsGrowthChangeHistoryModel:           umsmodel.NewUmsGrowthChangeHistoryModel(sqlConn),
		UmsIntegrationChangeHistoryModel:      umsmodel.NewUmsIntegrationChangeHistoryModel(sqlConn),
		UmsIntegrationConsumeSettingModel:     umsmodel.NewUmsIntegrationConsumeSettingModel(sqlConn),
		UmsMemberModel:                        umsmodel.NewUmsMemberModel(sqlConn),
		UmsMemberLevelModel:                   umsmodel.NewUmsMemberLevelModel(sqlConn),
		UmsMemberLoginLogModel:                umsmodel.NewUmsMemberLoginLogModel(sqlConn),
		UmsMemberMemberTagRelationModel:       umsmodel.NewUmsMemberMemberTagRelationModel(sqlConn),
		UmsMemberProductCategoryRelationModel: umsmodel.NewUmsMemberProductCategoryRelationModel(sqlConn),
		UmsMemberReceiveAddressModel:          umsmodel.NewUmsMemberReceiveAddressModel(sqlConn),
		UmsMemberRuleSettingModel:             umsmodel.NewUmsMemberRuleSettingModel(sqlConn),
		UmsMemberStatisticsInfoModel:          umsmodel.NewUmsMemberStatisticsInfoModel(sqlConn),
		UmsMemberTaskModel:                    umsmodel.NewUmsMemberTaskModel(sqlConn),
		UmsMemberTagModel:                     umsmodel.NewUmsMemberTagModel(sqlConn),
	}
}
