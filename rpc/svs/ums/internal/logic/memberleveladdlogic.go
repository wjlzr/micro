package logic

import (
	"context"
	"micro/rpc/model/umsmodel"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelAddLogic {
	return &MemberLevelAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelAddLogic) MemberLevelAdd(in *ums.MemberLevelAddReq) (*ums.MemberLevelAddResp, error) {
	_, err := l.svcCtx.UmsMemberLevelModel.Insert(umsmodel.UmsMemberLevel{
		Name:                  in.Name,
		GrowthPoint:           in.GrowthPoint,
		DefaultStatus:         in.DefaultStatus,
		FreeFreightPoint:      float64(in.FreeFreightPoint),
		CommentGrowthPoint:    in.CommentGrowthPoint,
		PriviledgeFreeFreight: in.PriviledgeFreeFreight,
		PriviledgeSignIn:      in.PriviledgeSignIn,
		PriviledgeComment:     in.PriviledgeComment,
		PriviledgePromotion:   in.PriviledgePromotion,
		PriviledgeMemberPrice: in.PriviledgeMemberPrice,
		PriviledgeBirthday:    in.PriviledgeBirthday,
		Note:                  in.Note,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberLevelAddResp{}, nil
}
