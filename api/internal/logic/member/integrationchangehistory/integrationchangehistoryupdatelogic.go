package logic

import (
	"context"
	"encoding/json"
	"micro/common/errorx"
	"micro/rpc/svs/ums/umsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationChangeHistoryUpdateLogic {
	return IntegrationChangeHistoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryUpdateLogic) IntegrationChangeHistoryUpdate(req types.UpdateIntegrationChangeHistoryReq) (*types.UpdateIntegrationChangeHistoryResp, error) {
	_, err := l.svcCtx.Ums.IntegrationChangeHistoryUpdate(l.ctx, &umsclient.IntegrationChangeHistoryUpdateReq{
		Id:          req.Id,
		MemberId:    req.MemberId,
		CreateTime:  req.CreateTime,
		ChangeType:  req.ChangeType,
		ChangeCount: req.ChangeCount,
		OperateMan:  req.OperateMan,
		OperateNote: req.OperateNote,
		SourceType:  req.SourceType,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新积分变化历史记录信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新积分变化历史记录失败")
	}

	return &types.UpdateIntegrationChangeHistoryResp{
		Code:    "000000",
		Message: "更新积分变化历史记录成功",
	}, nil
}
