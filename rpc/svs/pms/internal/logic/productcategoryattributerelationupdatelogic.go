package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryAttributeRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationUpdateLogic {
	return &ProductCategoryAttributeRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryAttributeRelationUpdateLogic) ProductCategoryAttributeRelationUpdate(in *pms.ProductCategoryAttributeRelationUpdateReq) (*pms.ProductCategoryAttributeRelationUpdateResp, error) {
	err := l.svcCtx.PmsProductCategoryAttributeRelationModel.Update(pmsmodel.PmsProductCategoryAttributeRelation{
		Id:                 in.Id,
		ProductCategoryId:  in.ProductCategoryId,
		ProductAttributeId: in.ProductAttributeId,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductCategoryAttributeRelationUpdateResp{}, nil
}
