package logic

import (
	"context"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberDeleteLogic {
	return &MemberDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberDeleteLogic) MemberDelete(in *ums.MemberDeleteReq) (*ums.MemberDeleteResp, error) {
	err := l.svcCtx.UmsMemberModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}
	return &ums.MemberDeleteResp{}, nil
}
