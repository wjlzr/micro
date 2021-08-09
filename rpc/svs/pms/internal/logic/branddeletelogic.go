package logic

import (
	"context"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type BrandDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandDeleteLogic {
	return &BrandDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BrandDeleteLogic) BrandDelete(in *pms.BrandDeleteReq) (*pms.BrandDeleteResp, error) {
	err := l.svcCtx.PmsBrandModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.BrandDeleteResp{}, nil
}
