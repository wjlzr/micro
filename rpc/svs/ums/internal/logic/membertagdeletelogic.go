package logic

import (
	"context"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagDeleteLogic {
	return &MemberTagDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagDeleteLogic) MemberTagDelete(in *ums.MemberTagDeleteReq) (*ums.MemberTagDeleteResp, error) {
	err := l.svcCtx.UmsMemberTagModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberTagDeleteResp{}, nil
}
