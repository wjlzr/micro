package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryAttributeRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationAddLogic {
	return &ProductCategoryAttributeRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryAttributeRelationAddLogic) ProductCategoryAttributeRelationAdd(in *pms.ProductCategoryAttributeRelationAddReq) (*pms.ProductCategoryAttributeRelationAddResp, error) {
	_, err := l.svcCtx.PmsProductCategoryAttributeRelationModel.Insert(pmsmodel.PmsProductCategoryAttributeRelation{
		ProductCategoryId:  in.ProductCategoryId,
		ProductAttributeId: in.ProductAttributeId,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductCategoryAttributeRelationAddResp{}, nil
}
