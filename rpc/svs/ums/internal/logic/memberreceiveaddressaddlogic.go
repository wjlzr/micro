package logic

import (
	"context"
	"micro/rpc/model/umsmodel"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberReceiveAddressAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressAddLogic {
	return &MemberReceiveAddressAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressAddLogic) MemberReceiveAddressAdd(in *ums.MemberReceiveAddressAddReq) (*ums.MemberReceiveAddressAddResp, error) {
	_, err := l.svcCtx.UmsMemberReceiveAddressModel.Insert(umsmodel.UmsMemberReceiveAddress{
		MemberId:      in.MemberId,
		Name:          in.Name,
		PhoneNumber:   in.PhoneNumber,
		DefaultStatus: in.DefaultStatus,
		PostCode:      in.PostCode,
		Province:      in.Province,
		City:          in.City,
		Region:        in.Region,
		DetailAddress: in.DetailAddress,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberReceiveAddressAddResp{}, nil
}
