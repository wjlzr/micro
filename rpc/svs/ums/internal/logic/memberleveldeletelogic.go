package logic

import (
	"context"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelDeleteLogic {
	return &MemberLevelDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelDeleteLogic) MemberLevelDelete(in *ums.MemberLevelDeleteReq) (*ums.MemberLevelDeleteResp, error) {
	err := l.svcCtx.UmsMemberLevelModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberLevelDeleteResp{}, nil
}
