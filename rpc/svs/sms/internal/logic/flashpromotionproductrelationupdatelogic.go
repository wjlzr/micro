package logic

import (
	"context"
	"micro/rpc/model/smsmodel"

	"micro/rpc/svs/sms/internal/svc"
	"micro/rpc/svs/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionProductRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionProductRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationUpdateLogic {
	return &FlashPromotionProductRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionProductRelationUpdateLogic) FlashPromotionProductRelationUpdate(in *sms.FlashPromotionProductRelationUpdateReq) (*sms.FlashPromotionProductRelationUpdateResp, error) {
	err := l.svcCtx.SmsFlashPromotionProductRelationModel.Update(smsmodel.SmsFlashPromotionProductRelation{
		Id:                      in.Id,
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

	return &sms.FlashPromotionProductRelationUpdateResp{}, nil
}
