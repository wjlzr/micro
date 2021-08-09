package logic

import (
	"context"
	"micro/common/errorx"
	"micro/api/internal/svc"
	"micro/api/internal/types"
	"micro/rpc/svs/oms/omsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompayAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressDeleteLogic {
	return CompayAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressDeleteLogic) CompayAddressDelete(req types.DeleteCompayAddressReq) (*types.DeleteCompayAddressResp, error) {
	_, err := l.svcCtx.Oms.CompanyAddressDelete(l.ctx, &omsclient.CompanyAddressDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除公司地址异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除公司地址失败")
	}

	return &types.DeleteCompayAddressResp{
		Code:    "000000",
		Message: "",
	}, nil
}
