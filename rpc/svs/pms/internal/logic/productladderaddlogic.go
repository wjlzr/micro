package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductLadderAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderAddLogic {
	return &ProductLadderAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLadderAddLogic) ProductLadderAdd(in *pms.ProductLadderAddReq) (*pms.ProductLadderAddResp, error) {
	_, err := l.svcCtx.PmsProductLadderModel.Insert(pmsmodel.PmsProductLadder{
		ProductId: in.ProductId,
		Count:     in.Count,
		Discount:  float64(in.Discount),
		Price:     float64(in.Price),
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductLadderAddResp{}, nil
}
