package logic

import (
	"context"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductLadderDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderDeleteLogic {
	return &ProductLadderDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLadderDeleteLogic) ProductLadderDelete(in *pms.ProductLadderDeleteReq) (*pms.ProductLadderDeleteResp, error) {
	err := l.svcCtx.PmsProductLadderModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductLadderDeleteResp{}, nil
}
