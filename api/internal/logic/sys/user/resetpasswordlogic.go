package logic

import (
	"context"
	"micro/rpc/svs/sys/sysclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReSetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReSetPasswordLogic {
	return ReSetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReSetPasswordLogic) ReSetPassword(req types.ReSetPasswordReq) (*types.ReSetPasswordResp, error) {
	_, _ = l.svcCtx.Sys.ReSetPassword(l.ctx, &sysclient.ReSetPasswordReq{
		Id:           req.Id,
		LastUpdateBy: "admin",
	})

	return &types.ReSetPasswordResp{
		Code:    "000000",
		Message: "",
	}, nil
}
