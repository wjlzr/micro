package logic

import (
	"context"
	"micro/common/errorx"
	"micro/rpc/svs/sms/smsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponHistoryDeleteLogic {
	return CouponHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponHistoryDeleteLogic) CouponHistoryDelete(req types.DeleteCouponHistoryReq) (*types.DeleteCouponHistoryResp, error) {
	_, err := l.svcCtx.Sms.CouponHistoryDelete(l.ctx, &smsclient.CouponHistoryDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除优惠券使用记录异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除优惠券使用记录失败")
	}
	return &types.DeleteCouponHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
