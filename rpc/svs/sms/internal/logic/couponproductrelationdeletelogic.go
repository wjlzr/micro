package logic

import (
	"context"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponProductRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductRelationDeleteLogic {
	return &CouponProductRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductRelationDeleteLogic) CouponProductRelationDelete(in *sms.CouponProductRelationDeleteReq) (*sms.CouponProductRelationDeleteResp, error) {
	err := l.svcCtx.SmsCouponProductRelationModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.CouponProductRelationDeleteResp{}, nil
}
