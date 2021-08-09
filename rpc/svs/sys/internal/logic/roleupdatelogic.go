package logic

import (
	"context"
	"micro/rpc/model/sysmodel"
	"time"

	"micro/rpc/svs/sys/internal/svc"
	"micro/rpc/svs/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *sys.RoleUpdateReq) (*sys.RoleUpdateResp, error) {
	_ = l.svcCtx.RoleModel.Update(sysmodel.SysRole{
		Id:             in.Id,
		Name:           in.Name,
		Remark:         in.Remark,
		CreateBy:       "admin",
		CreateTime:     time.Time{},
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
		Status:         in.Status,
	})

	return &sys.RoleUpdateResp{}, nil
}
