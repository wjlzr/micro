package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductFullReductionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionUpdateLogic {
	return &ProductFullReductionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFullReductionUpdateLogic) ProductFullReductionUpdate(in *pms.ProductFullReductionUpdateReq) (*pms.ProductFullReductionUpdateResp, error) {
	err := l.svcCtx.PmsProductFullReductionModel.Update(pmsmodel.PmsProductFullReduction{
		Id:          in.Id,
		ProductId:   in.ProductId,
		FullPrice:   float64(in.FullPrice),
		ReducePrice: float64(in.ReducePrice),
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductFullReductionUpdateResp{}, nil
}
