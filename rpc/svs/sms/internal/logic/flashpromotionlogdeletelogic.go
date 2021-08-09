package logic

import (
	"context"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionLogDeleteLogic {
	return &FlashPromotionLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionLogDeleteLogic) FlashPromotionLogDelete(in *sms.FlashPromotionLogDeleteReq) (*sms.FlashPromotionLogDeleteResp, error) {
	err := l.svcCtx.SmsFlashPromotionLogModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionLogDeleteResp{}, nil
}
