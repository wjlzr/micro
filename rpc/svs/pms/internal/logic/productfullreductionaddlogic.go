package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductFullReductionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionAddLogic {
	return &ProductFullReductionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFullReductionAddLogic) ProductFullReductionAdd(in *pms.ProductFullReductionAddReq) (*pms.ProductFullReductionAddResp, error) {
	_, err := l.svcCtx.PmsProductFullReductionModel.Insert(pmsmodel.PmsProductFullReduction{
		ProductId:   in.ProductId,
		FullPrice:   float64(in.FullPrice),
		ReducePrice: float64(in.ReducePrice),
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductFullReductionAddResp{}, nil
}
