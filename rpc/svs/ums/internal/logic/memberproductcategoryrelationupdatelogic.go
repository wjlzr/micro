package logic

import (
	"context"
	"micro/rpc/model/umsmodel"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberProductCategoryRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationUpdateLogic {
	return &MemberProductCategoryRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationUpdateLogic) MemberProductCategoryRelationUpdate(in *ums.MemberProductCategoryRelationUpdateReq) (*ums.MemberProductCategoryRelationUpdateResp, error) {
	err := l.svcCtx.UmsMemberProductCategoryRelationModel.Update(umsmodel.UmsMemberProductCategoryRelation{
		Id:                in.Id,
		MemberId:          in.MemberId,
		ProductCategoryId: in.ProductCategoryId,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberProductCategoryRelationUpdateResp{}, nil
}
