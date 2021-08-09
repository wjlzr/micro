package logic

import (
	"context"
	"micro/rpc/model/smsmodel"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponProductCategoryRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductCategoryRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductCategoryRelationUpdateLogic {
	return &CouponProductCategoryRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductCategoryRelationUpdateLogic) CouponProductCategoryRelationUpdate(in *sms.CouponProductCategoryRelationUpdateReq) (*sms.CouponProductCategoryRelationUpdateResp, error) {
	err := l.svcCtx.SmsCouponProductCategoryRelationModel.Update(smsmodel.SmsCouponProductCategoryRelation{
		Id:                  in.Id,
		CouponId:            in.CouponId,
		ProductCategoryId:   in.ProductCategoryId,
		ProductCategoryName: in.ProductCategoryName,
		ParentCategoryName:  in.ParentCategoryName,
	})
	if err != nil {
		return nil, err
	}

	return &sms.CouponProductCategoryRelationUpdateResp{}, nil
}
