package logic

import (
	"context"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeNewProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeNewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeNewProductDeleteLogic {
	return &HomeNewProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeNewProductDeleteLogic) HomeNewProductDelete(in *sms.HomeNewProductDeleteReq) (*sms.HomeNewProductDeleteResp, error) {
	err := l.svcCtx.SmsHomeNewProductModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.HomeNewProductDeleteResp{}, nil
}
