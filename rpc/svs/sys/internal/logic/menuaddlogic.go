package logic

import (
	"context"
	"micro/rpc/model/sysmodel"
	"time"

	"micro/rpc/svs/sys/internal/svc"
	"micro/rpc/svs/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuAddLogic) MenuAdd(in *sys.MenuAddReq) (*sys.MenuAddResp, error) {
	_, err := l.svcCtx.MenuModel.Insert(sysmodel.SysMenu{
		Id:             0,
		Name:           in.Name,
		ParentId:       in.ParentId,
		Url:            in.Url,
		Perms:          in.Perms,
		Type:           in.Type,
		Icon:           in.Icon,
		OrderNum:       in.OrderNum,
		CreateBy:       in.CreateBy,
		CreateTime:     time.Time{},
		LastUpdateBy:   in.CreateBy,
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

	return &sys.MenuAddResp{}, nil
}
