package logic

import (
	"context"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryDeleteLogic {
	return &IntegrationChangeHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationChangeHistoryDeleteLogic) IntegrationChangeHistoryDelete(in *ums.IntegrationChangeHistoryDeleteReq) (*ums.IntegrationChangeHistoryDeleteResp, error) {
	err := l.svcCtx.UmsIntegrationChangeHistoryModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.IntegrationChangeHistoryDeleteResp{}, nil
}
