package logic

import (
	"context"
	"micro/rpc/model/sysmodel"
	"time"

	"micro/rpc/svs/sys/internal/svc"
	"micro/rpc/svs/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictUpdateLogic {
	return &DictUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictUpdateLogic) DictUpdate(in *sys.DictUpdateReq) (*sys.DictUpdateResp, error) {
	err := l.svcCtx.DictModel.Update(sysmodel.SysDict{
		Id:             in.Id,
		Value:          in.Value,
		Label:          in.Label,
		Type:           in.Type,
		Description:    in.Description,
		Sort:           float64(in.Sort),
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
		Remarks:        in.Remarks,
		DelFlag:        0,
	})

	if err != nil {
		return nil, err
	}

	return &sys.DictUpdateResp{}, nil
}
