package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeAddLogic {
	return &ProductAttributeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeAddLogic) ProductAttributeAdd(in *pms.ProductAttributeAddReq) (*pms.ProductAttributeAddResp, error) {
	_, err := l.svcCtx.PmsProductAttributeModel.Insert(pmsmodel.PmsProductAttribute{
		ProductAttributeCategoryId: in.ProductAttributeCategoryId,
		Name:                       in.Name,
		SelectType:                 in.SelectType,
		InputType:                  in.InputType,
		InputList:                  in.InputList,
		Sort:                       in.Sort,
		FilterType:                 in.FilterType,
		SearchType:                 in.SearchType,
		RelatedStatus:              in.RelatedStatus,
		HandAddStatus:              in.HandAddStatus,
		Type:                       in.Type,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeAddResp{}, nil
}
