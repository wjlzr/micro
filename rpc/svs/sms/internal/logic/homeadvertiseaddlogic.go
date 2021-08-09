package logic

import (
	"context"
	"micro/rpc/model/smsmodel"
	"time"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeAdvertiseAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseAddLogic {
	return &HomeAdvertiseAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeAdvertiseAddLogic) HomeAdvertiseAdd(in *sms.HomeAdvertiseAddReq) (*sms.HomeAdvertiseAddResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	_, err := l.svcCtx.SmsHomeAdvertiseModel.Insert(smsmodel.SmsHomeAdvertise{
		Name:       in.Name,
		Type:       in.Type,
		Pic:        in.Pic,
		StartTime:  StartTime,
		EndTime:    EndTime,
		Status:     in.Status,
		ClickCount: in.ClickCount,
		OrderCount: in.OrderCount,
		Url:        in.Url,
		Note:       in.Note,
		Sort:       in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &sms.HomeAdvertiseAddResp{}, nil
}
