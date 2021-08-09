package logic

import (
	"context"
	"micro/common/errorx"
	"micro/rpc/svs/ums/umsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationConsumeSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationConsumeSettingDeleteLogic {
	return IntegrationConsumeSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingDeleteLogic) IntegrationConsumeSettingDelete(req types.DeleteIntegrationConsumeSettingReq) (*types.DeleteIntegrationConsumeSettingResp, error) {
	_, err := l.svcCtx.Ums.IntegrationConsumeSettingDelete(l.ctx, &umsclient.IntegrationConsumeSettingDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除积分消费设置异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除积分消费设置失败")
	}
	return &types.DeleteIntegrationConsumeSettingResp{
		Code:    "000000",
		Message: "删除积分消费设置成功",
	}, nil
}
