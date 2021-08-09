package logic

import (
	"context"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryDeleteLogic {
	return &CouponHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponHistoryDeleteLogic) CouponHistoryDelete(in *sms.CouponHistoryDeleteReq) (*sms.CouponHistoryDeleteResp, error) {
	err := l.svcCtx.SmsCouponHistoryModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.CouponHistoryDeleteResp{}, nil
}
