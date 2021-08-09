package logic

import (
	"context"

	"micro/rpc/svs/oms/internal/svc"
	"micro/rpc/svs/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompanyAddressDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddressDeleteLogic {
	return &CompanyAddressDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompanyAddressDeleteLogic) CompanyAddressDelete(in *oms.CompanyAddressDeleteReq) (*oms.CompanyAddressDeleteResp, error) {
	err := l.svcCtx.OmsCompanyAddressModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &oms.CompanyAddressDeleteResp{}, nil
}
