package logic

import (
	"context"
	"micro/rpc/model/sysmodel"
	"time"

	"micro/rpc/svs/sys/internal/svc"
	"micro/rpc/svs/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuUpdateLogic) MenuUpdate(in *sys.MenuUpdateReq) (*sys.MenuUpdateResp, error) {
	err := l.svcCtx.MenuModel.Update(sysmodel.SysMenu{
		Id:             in.Id,
		Name:           in.Name,
		ParentId:       in.ParentId,
		Url:            in.Url,
		Perms:          in.Perms,
		Type:           in.Type,
		Icon:           in.Icon,
		OrderNum:       in.OrderNum,
		CreateTime:     time.Time{},
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
		VuePath:        in.VuePath,
		VueComponent:   in.VueComponent,
		VueIcon:        in.VueIcon,
		VueRedirect:    in.VueRedirect,
	})
	//count, _ := l.svcCtx.UserModel.Count()

	if err != nil {
		return nil, err
	}

	return &sys.MenuUpdateResp{}, nil
}
