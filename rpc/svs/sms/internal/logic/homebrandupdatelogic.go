package logic

import (
	"context"
	"micro/rpc/model/smsmodel"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandUpdateLogic {
	return &HomeBrandUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeBrandUpdateLogic) HomeBrandUpdate(in *sms.HomeBrandUpdateReq) (*sms.HomeBrandUpdateResp, error) {
	err := l.svcCtx.SmsHomeBrandModel.Update(smsmodel.SmsHomeBrand{
		Id:              in.Id,
		BrandId:         in.BrandId,
		BrandName:       in.BrandName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &sms.HomeBrandUpdateResp{}, nil
}
