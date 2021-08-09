package logic

import (
	"context"
	"micro/rpc/model/smsmodel"
	"time"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponAddLogic {
	return &CouponAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponAddLogic) CouponAdd(in *sms.CouponAddReq) (*sms.CouponAddResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	EnableTime, _ := time.Parse("2006-01-02 15:04:05", in.EnableTime)
	_, err := l.svcCtx.SmsCouponModel.Insert(smsmodel.SmsCoupon{
		Type:         in.Type,
		Name:         in.Name,
		Platform:     in.Platform,
		Count:        in.Count,
		Amount:       in.Amount,
		PerLimit:     in.PerLimit,
		MinPoint:     in.MinPoint,
		StartTime:    StartTime,
		EndTime:      EndTime,
		UseType:      in.UseType,
		Note:         in.Note,
		PublishCount: in.PublishCount,
		UseCount:     in.UseCount,
		ReceiveCount: in.ReceiveCount,
		EnableTime:   EnableTime,
		Code:         in.Code,
		MemberLevel:  in.MemberLevel,
	})
	if err != nil {
		return nil, err
	}

	return &sms.CouponAddResp{}, nil
}
