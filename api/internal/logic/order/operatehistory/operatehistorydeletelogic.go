package logic

import (
	"context"
	"micro/common/errorx"
	"micro/api/internal/svc"
	"micro/api/internal/types"
	"micro/rpc/svs/oms/omsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type OperateHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryDeleteLogic {
	return OperateHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryDeleteLogic) OperateHistoryDelete(req types.DeleteOperateHistoryReq) (*types.DeleteOperateHistoryResp, error) {
	_, err := l.svcCtx.Oms.OrderOperateHistoryDelete(l.ctx, &omsclient.OrderOperateHistoryDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除订单操作历史异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除订单操作历史签失败")
	}

	return &types.DeleteOperateHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
