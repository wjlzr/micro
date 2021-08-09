package logic

import (
	"context"
	"micro/rpc/model/smsmodel"
	"time"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionSessionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionSessionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionUpdateLogic {
	return &FlashPromotionSessionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionSessionUpdateLogic) FlashPromotionSessionUpdate(in *sms.FlashPromotionSessionUpdateReq) (*sms.FlashPromotionSessionUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.SmsFlashPromotionSessionModel.Update(smsmodel.SmsFlashPromotionSession{
		Id:         in.Id,
		Name:       in.Name,
		StartTime:  in.StartTime,
		EndTime:    in.EndTime,
		Status:     in.Status,
		CreateTime: CreateTime,
	})
	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionSessionUpdateResp{}, nil
}
