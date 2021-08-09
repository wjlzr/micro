package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductLadderUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderUpdateLogic {
	return &ProductLadderUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLadderUpdateLogic) ProductLadderUpdate(in *pms.ProductLadderUpdateReq) (*pms.ProductLadderUpdateResp, error) {
	err := l.svcCtx.PmsProductLadderModel.Update(pmsmodel.PmsProductLadder{
		Id:        in.Id,
		ProductId: in.ProductId,
		Count:     in.Count,
		Discount:  float64(in.Discount),
		Price:     float64(in.Price),
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductLadderUpdateResp{}, nil
}
