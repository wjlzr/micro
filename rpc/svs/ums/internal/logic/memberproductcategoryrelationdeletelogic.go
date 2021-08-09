package logic

import (
	"context"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberProductCategoryRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationDeleteLogic {
	return &MemberProductCategoryRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationDeleteLogic) MemberProductCategoryRelationDelete(in *ums.MemberProductCategoryRelationDeleteReq) (*ums.MemberProductCategoryRelationDeleteResp, error) {
	err := l.svcCtx.UmsMemberProductCategoryRelationModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberProductCategoryRelationDeleteResp{}, nil
}
