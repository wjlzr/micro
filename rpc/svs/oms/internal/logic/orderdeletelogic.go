package logic

import (
	"context"

	"micro/rpc/svs/oms/internal/svc"
	"micro/rpc/svs/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteLogic {
	return &OrderDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderDeleteLogic) OrderDelete(in *oms.OrderDeleteReq) (*oms.OrderDeleteResp, error) {
	err := l.svcCtx.OmsOrderModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &oms.OrderDeleteResp{}, nil
}
