package logic

import (
	"context"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberPriceDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPriceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceDeleteLogic {
	return &MemberPriceDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPriceDeleteLogic) MemberPriceDelete(in *pms.MemberPriceDeleteReq) (*pms.MemberPriceDeleteResp, error) {
	err := l.svcCtx.PmsMemberPriceModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.MemberPriceDeleteResp{}, nil
}
