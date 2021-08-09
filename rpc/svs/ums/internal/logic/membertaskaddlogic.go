package logic

import (
	"context"
	"micro/rpc/model/umsmodel"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskAddLogic {
	return &MemberTaskAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskAddLogic) MemberTaskAdd(in *ums.MemberTaskAddReq) (*ums.MemberTaskAddResp, error) {
	_, err := l.svcCtx.UmsMemberTaskModel.Insert(umsmodel.UmsMemberTask{
		Name:         in.Name,
		Growth:       in.Growth,
		Intergration: in.Intergration,
		Type:         in.Type,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberTaskAddResp{}, nil
}
