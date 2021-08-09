package logic

import (
	"context"

	"micro/rpc/svs/oms/internal/svc"
	"micro/rpc/svs/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderOperateHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderOperateHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderOperateHistoryDeleteLogic {
	return &OrderOperateHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderOperateHistoryDeleteLogic) OrderOperateHistoryDelete(in *oms.OrderOperateHistoryDeleteReq) (*oms.OrderOperateHistoryDeleteResp, error) {
	err := l.svcCtx.OmsOrderOperateHistoryModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &oms.OrderOperateHistoryDeleteResp{}, nil
}
