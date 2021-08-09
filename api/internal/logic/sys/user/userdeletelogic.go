package logic

import (
	"context"
	"micro/common/errorx"
	"micro/api/internal/svc"
	"micro/api/internal/types"
	"micro/rpc/svs/sys/sysclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserDeleteLogic {
	return UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req types.DeleteUserReq) (*types.DeleteUserResp, error) {

	_, err := l.svcCtx.Sys.UserDelete(l.ctx, &sysclient.UserDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据userId: %d,删除用户异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除用户失败")
	}

	return &types.DeleteUserResp{
		Code:    "000000",
		Message: "删除用户成功",
	}, nil
}
