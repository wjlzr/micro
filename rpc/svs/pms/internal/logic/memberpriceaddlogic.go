package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberPriceAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPriceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceAddLogic {
	return &MemberPriceAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPriceAddLogic) MemberPriceAdd(in *pms.MemberPriceAddReq) (*pms.MemberPriceAddResp, error) {
	_, err := l.svcCtx.PmsMemberPriceModel.Insert(pmsmodel.PmsMemberPrice{
		ProductId:       in.ProductId,
		MemberLevelId:   in.MemberLevelId,
		MemberPrice:     float64(in.MemberPrice),
		MemberLevelName: in.MemberLevelName,
	})
	if err != nil {
		return nil, err
	}

	return &pms.MemberPriceAddResp{}, nil
}
