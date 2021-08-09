package logic

import (
	"context"
	"micro/common/errorx"
	"micro/api/internal/svc"
	"micro/api/internal/types"
	"micro/rpc/svs/oms/omsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnResonDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnResonDeleteLogic {
	return ReturnResonDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonDeleteLogic) ReturnResonDelete(req types.DeleteReturnResonReq) (*types.DeleteReturnResonResp, error) {
	_, err := l.svcCtx.Oms.OrderReturnReasonDelete(l.ctx, &omsclient.OrderReturnReasonDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除退货原因异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除退货原因失败")
	}
	return &types.DeleteReturnResonResp{
		Code:    "000000",
		Message: "",
	}, nil
}
