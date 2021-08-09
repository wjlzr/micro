package logic

import (
	"context"
	"micro/rpc/model/sysmodel"
	"time"

	"micro/rpc/svs/sys/internal/svc"
	"micro/rpc/svs/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *sys.UserUpdateReq) (*sys.UserUpdateResp, error) {

	_ = l.svcCtx.UserModel.Update(sysmodel.SysUser{
		Id:             in.Id,
		Name:           in.Name,
		NickName:       in.NickName,
		Avatar:         in.Avatar,
		Email:          in.Email,
		Mobile:         in.Mobile,
		DeptId:         in.DeptId,
		LastUpdateBy:   in.LastUpdateBy,
		Status:         in.Status,
		LastUpdateTime: time.Now(),
		JobId:          in.JobId,
	})

	_ = l.svcCtx.UserRoleModel.Delete(in.Id)

	_, _ = l.svcCtx.UserRoleModel.Insert(sysmodel.SysUserRole{
		UserId:         in.Id,
		RoleId:         in.RoleId,
		CreateBy:       "admin",
		CreateTime:     time.Now(),
		LastUpdateBy:   "admin",
		LastUpdateTime: time.Now(),
	})

	return &sys.UserUpdateResp{}, nil
}
