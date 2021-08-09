package logic

import (
	"context"
	"micro/rpc/model/smsmodel"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponProductCategoryRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductCategoryRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductCategoryRelationAddLogic {
	return &CouponProductCategoryRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductCategoryRelationAddLogic) CouponProductCategoryRelationAdd(in *sms.CouponProductCategoryRelationAddReq) (*sms.CouponProductCategoryRelationAddResp, error) {
	_, err := l.svcCtx.SmsCouponProductCategoryRelationModel.Insert(smsmodel.SmsCouponProductCategoryRelation{
		CouponId:            in.CouponId,
		ProductCategoryId:   in.ProductCategoryId,
		ProductCategoryName: in.ProductCategoryName,
		ParentCategoryName:  in.ParentCategoryName,
	})
	if err != nil {
		return nil, err
	}

	return &sms.CouponProductCategoryRelationAddResp{}, nil
}
