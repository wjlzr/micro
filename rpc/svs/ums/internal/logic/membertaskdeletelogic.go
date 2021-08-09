package logic

import (
	"context"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskDeleteLogic {
	return &MemberTaskDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskDeleteLogic) MemberTaskDelete(in *ums.MemberTaskDeleteReq) (*ums.MemberTaskDeleteResp, error) {
	err := l.svcCtx.UmsMemberTaskModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberTaskDeleteResp{}, nil
}
