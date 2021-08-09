package logic

import (
	"context"
	"micro/common/errorx"
	"micro/rpc/svs/pms/pmsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCommentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCommentDeleteLogic {
	return ProductCommentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentDeleteLogic) ProductCommentDelete(req types.DeleteProductCommentReq) (*types.DeleteProductCommentResp, error) {
	_, err := l.svcCtx.Pms.CommentDelete(l.ctx, &pmsclient.CommentDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除商品评价异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除商品评价失败")
	}

	return &types.DeleteProductCommentResp{
		Code:    "000000",
		Message: "",
	}, nil
}
