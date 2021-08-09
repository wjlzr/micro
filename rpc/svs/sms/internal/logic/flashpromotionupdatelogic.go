package logic

import (
	"context"
	"micro/rpc/model/smsmodel"
	"time"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionUpdateLogic {
	return &FlashPromotionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionUpdateLogic) FlashPromotionUpdate(in *sms.FlashPromotionUpdateReq) (*sms.FlashPromotionUpdateResp, error) {
	StartDate, _ := time.Parse("2006-01-02 15:04:05", in.StartDate)
	EndDate, _ := time.Parse("2006-01-02 15:04:05", in.EndDate)
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.SmsFlashPromotionModel.Update(smsmodel.SmsFlashPromotion{
		Id:         in.Id,
		Title:      in.Title,
		StartDate:  StartDate,
		EndDate:    EndDate,
		Status:     in.Status,
		CreateTime: CreateTime,
	})
	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionUpdateResp{}, nil
}
