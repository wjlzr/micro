package logic

import (
	"context"
	"encoding/json"
	"micro/common/errorx"
	"micro/rpc/svs/pms/pmsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductBrandUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductBrandUpdateLogic {
	return ProductBrandUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductBrandUpdateLogic) ProductBrandUpdate(req types.UpdateProductBrandReq) (*types.UpdateProductBrandResp, error) {
	_, err := l.svcCtx.Pms.BrandUpdate(l.ctx, &pmsclient.BrandUpdateReq{
		Id:                  req.Id,
		Name:                req.Name,
		FirstLetter:         req.FirstLetter,
		Sort:                req.Sort,
		FactoryStatus:       req.FactoryStatus,
		ShowStatus:          req.ShowStatus,
		ProductCount:        req.ProductCount,
		ProductCommentCount: req.ProductCommentCount,
		Logo:                req.Logo,
		BigPic:              req.BigPic,
		BrandStory:          req.BrandStory,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新商品品牌信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新商品品牌失败")
	}

	return &types.UpdateProductBrandResp{
		Code:    "000000",
		Message: "更新商品品牌成功",
	}, nil
}
