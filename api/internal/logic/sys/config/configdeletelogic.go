package logic

import (
	"context"
	"micro/common/errorx"
	"micro/rpc/svs/sys/sysclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigDeleteLogic {
	return ConfigDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigDeleteLogic) ConfigDelete(req types.DeleteConfigReq) (*types.DeleteConfigResp, error) {
	_, err := l.svcCtx.Sys.DeptDelete(l.ctx, &sysclient.DeptDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("删除参数配置失败")
	}

	return &types.DeleteConfigResp{}, nil
}
