package logic

import (
	"context"
	"micro/common/errorx"
	"micro/rpc/svs/sms/smsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandDeleteLogic {
	return HomeBrandDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandDeleteLogic) HomeBrandDelete(req types.DeleteHomeBrandReq) (*types.DeleteHomeBrandResp, error) {
	_, err := l.svcCtx.Sms.HomeBrandDelete(l.ctx, &smsclient.HomeBrandDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除首页品牌异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除首页品牌失败")
	}

	return &types.DeleteHomeBrandResp{
		Code:    "000000",
		Message: "",
	}, nil
}
