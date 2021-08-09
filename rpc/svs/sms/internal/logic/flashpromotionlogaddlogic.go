package logic

import (
	"context"
	"micro/rpc/model/smsmodel"
	"time"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionLogAddLogic {
	return &FlashPromotionLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionLogAddLogic) FlashPromotionLogAdd(in *sms.FlashPromotionLogAddReq) (*sms.FlashPromotionLogAddResp, error) {
	SubscribeTime, _ := time.Parse("2006-01-02 15:04:05", in.SubscribeTime)
	SendTime, _ := time.Parse("2006-01-02 15:04:05", in.SendTime)
	_, err := l.svcCtx.SmsFlashPromotionLogModel.Insert(smsmodel.SmsFlashPromotionLog{
		MemberId:      in.MemberId,
		ProductId:     in.ProductId,
		MemberPhone:   in.MemberPhone,
		ProductName:   in.ProductName,
		SubscribeTime: SubscribeTime,
		SendTime:      SendTime,
	})
	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionLogAddResp{}, nil
}
