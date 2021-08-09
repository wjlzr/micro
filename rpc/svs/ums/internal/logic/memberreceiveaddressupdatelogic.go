package logic

import (
	"context"
	"micro/rpc/model/umsmodel"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberReceiveAddressUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressUpdateLogic {
	return &MemberReceiveAddressUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressUpdateLogic) MemberReceiveAddressUpdate(in *ums.MemberReceiveAddressUpdateReq) (*ums.MemberReceiveAddressUpdateResp, error) {
	err := l.svcCtx.UmsMemberReceiveAddressModel.Update(umsmodel.UmsMemberReceiveAddress{
		Id:            in.Id,
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

	return &ums.MemberReceiveAddressUpdateResp{}, nil
}
