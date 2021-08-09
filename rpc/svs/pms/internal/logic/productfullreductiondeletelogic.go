package logic

import (
	"context"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductFullReductionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionDeleteLogic {
	return &ProductFullReductionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFullReductionDeleteLogic) ProductFullReductionDelete(in *pms.ProductFullReductionDeleteReq) (*pms.ProductFullReductionDeleteResp, error) {
	err := l.svcCtx.PmsProductFullReductionModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductFullReductionDeleteResp{}, nil
}
