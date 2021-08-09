package logic

import (
	"context"
	"micro/common/errorx"
	"micro/rpc/svs/pms/pmsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SkuStockDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SkuStockDeleteLogic {
	return SkuStockDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuStockDeleteLogic) SkuStockDelete(req types.DeleteSkuStockReq) (*types.DeleteSkuStockResp, error) {
	_, err := l.svcCtx.Pms.SkuStockDelete(l.ctx, &pmsclient.SkuStockDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除库存异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除库存失败")
	}

	return &types.DeleteSkuStockResp{
		Code:    "000000",
		Message: "",
	}, nil
}
