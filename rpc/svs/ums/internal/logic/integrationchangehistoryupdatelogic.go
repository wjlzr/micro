package logic

import (
	"context"
	"micro/rpc/model/umsmodel"
	"time"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryUpdateLogic {
	return &IntegrationChangeHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationChangeHistoryUpdateLogic) IntegrationChangeHistoryUpdate(in *ums.IntegrationChangeHistoryUpdateReq) (*ums.IntegrationChangeHistoryUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.UmsIntegrationChangeHistoryModel.Update(umsmodel.UmsIntegrationChangeHistory{
		Id:          in.Id,
		MemberId:    in.MemberId,
		CreateTime:  CreateTime,
		ChangeType:  in.ChangeType,
		ChangeCount: in.ChangeCount,
		OperateMan:  in.OperateMan,
		OperateNote: in.OperateNote,
		SourceType:  in.SourceType,
	})
	if err != nil {
		return nil, err
	}

	return &ums.IntegrationChangeHistoryUpdateResp{}, nil
}
