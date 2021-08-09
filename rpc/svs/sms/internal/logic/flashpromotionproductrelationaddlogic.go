package logic

import (
	"context"
	"micro/rpc/model/smsmodel"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionProductRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionProductRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationAddLogic {
	return &FlashPromotionProductRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionProductRelationAddLogic) FlashPromotionProductRelationAdd(in *sms.FlashPromotionProductRelationAddReq) (*sms.FlashPromotionProductRelationAddResp, error) {
	_, err := l.svcCtx.SmsFlashPromotionProductRelationModel.Insert(smsmodel.SmsFlashPromotionProductRelation{
		FlashPromotionId:        in.FlashPromotionId,
		FlashPromotionSessionId: in.FlashPromotionSessionId,
		ProductId:               in.ProductId,
		FlashPromotionPrice:     float64(in.FlashPromotionPrice),
		FlashPromotionCount:     in.FlashPromotionCount,
		FlashPromotionLimit:     in.FlashPromotionLimit,
		Sort:                    in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionProductRelationAddResp{}, nil
}
