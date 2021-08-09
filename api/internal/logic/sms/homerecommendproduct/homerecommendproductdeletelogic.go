package logic

import (
	"context"
	"micro/common/errorx"
	"micro/rpc/svs/sms/smsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendProductDeleteLogic {
	return HomeRecommendProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendProductDeleteLogic) HomeRecommendProductDelete(req types.DeleteHomeRecommendProductReq) (*types.DeleteHomeRecommendProductResp, error) {
	_, err := l.svcCtx.Sms.HomeRecommendProductDelete(l.ctx, &smsclient.HomeRecommendProductDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除人气推荐商品异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除人气推荐商品失败")
	}
	return &types.DeleteHomeRecommendProductResp{
		Code:    "000000",
		Message: "",
	}, nil
}
