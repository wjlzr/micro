package logic

import (
	"context"
	"micro/rpc/model/umsmodel"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberMemberTagRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberMemberTagRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberMemberTagRelationUpdateLogic {
	return &MemberMemberTagRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberMemberTagRelationUpdateLogic) MemberMemberTagRelationUpdate(in *ums.MemberMemberTagRelationUpdateReq) (*ums.MemberMemberTagRelationUpdateResp, error) {
	err := l.svcCtx.UmsMemberMemberTagRelationModel.Update(umsmodel.UmsMemberMemberTagRelation{
		Id:       in.Id,
		MemberId: in.MemberId,
		TagId:    in.TagId,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberMemberTagRelationUpdateResp{}, nil
}
