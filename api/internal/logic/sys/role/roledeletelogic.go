package logic

import (
	"context"
	"micro/common/errorx"
	"micro/rpc/svs/sys/sysclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleDeleteLogic {
	return RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req types.DeleteRoleReq) (*types.DeleteRoleResp, error) {
	_, err := l.svcCtx.Sys.RoleDelete(l.ctx, &sysclient.RoleDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据roleId: %d,删除角色异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除角色失败")
	}

	return &types.DeleteRoleResp{
		Code:    "000000",
		Message: "删除角色成功",
	}, nil
}
