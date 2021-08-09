package logic

import (
	"context"
	"encoding/json"
	"micro/common/errorx"
	"micro/rpc/svs/sms/smsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionSessionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionAddLogic {
	return FlashPromotionSessionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionAddLogic) FlashPromotionSessionAdd(req types.AddFlashPromotionSessionReq) (*types.AddFlashPromotionSessionResp, error) {
	_, err := l.svcCtx.Sms.FlashPromotionSessionAdd(l.ctx, &smsclient.FlashPromotionSessionAddReq{
		Name:       req.Name,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Status:     req.Status,
		CreateTime: req.CreateTime,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加限时购场次信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加限时购场次表失败")
	}

	return &types.AddFlashPromotionSessionResp{
		Code:    "000000",
		Message: "添加限时购场次表成功",
	}, nil
}
