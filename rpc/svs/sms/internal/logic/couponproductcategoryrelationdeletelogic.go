package logic

import (
	"context"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponProductCategoryRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductCategoryRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductCategoryRelationDeleteLogic {
	return &CouponProductCategoryRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductCategoryRelationDeleteLogic) CouponProductCategoryRelationDelete(in *sms.CouponProductCategoryRelationDeleteReq) (*sms.CouponProductCategoryRelationDeleteResp, error) {
	err := l.svcCtx.SmsCouponProductCategoryRelationModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.CouponProductCategoryRelationDeleteResp{}, nil
}
