package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeCategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryAddLogic {
	return &ProductAttributeCategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeCategoryAddLogic) ProductAttributeCategoryAdd(in *pms.ProductAttributeCategoryAddReq) (*pms.ProductAttributeCategoryAddResp, error) {
	_, err := l.svcCtx.PmsProductAttributeCategoryModel.Insert(pmsmodel.PmsProductAttributeCategory{
		Name:           in.Name,
		AttributeCount: in.AttributeCount,
		ParamCount:     in.ParamCount,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeCategoryAddResp{}, nil
}
