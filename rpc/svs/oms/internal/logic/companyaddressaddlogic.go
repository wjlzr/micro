package logic

import (
	"context"
	"micro/rpc/model/omsmodel"
	"micro/rpc/svs/oms/internal/svc"
	"micro/rpc/svs/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompanyAddressAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddressAddLogic {
	return &CompanyAddressAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompanyAddressAddLogic) CompanyAddressAdd(in *oms.CompanyAddressAddReq) (*oms.CompanyAddressAddResp, error) {
	_, err := l.svcCtx.OmsCompanyAddressModel.Insert(omsmodel.OmsCompanyAddress{
		AddressName:   in.AddressName,
		SendStatus:    in.SendStatus,
		ReceiveStatus: in.ReceiveStatus,
		Name:          in.Name,
		Phone:         in.Phone,
		Province:      in.Province,
		City:          in.City,
		Region:        in.Region,
		DetailAddress: in.DetailAddress,
	})

	if err != nil {
		return nil, err
	}

	return &oms.CompanyAddressAddResp{}, nil
}
