package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"
	"time"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductVertifyRecordAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordAddLogic {
	return &ProductVertifyRecordAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVertifyRecordAddLogic) ProductVertifyRecordAdd(in *pms.ProductVertifyRecordAddReq) (*pms.ProductVertifyRecordAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.PmsProductVertifyRecordModel.Insert(pmsmodel.PmsProductVertifyRecord{
		ProductId:  in.ProductId,
		CreateTime: CreateTime,
		VertifyMan: in.VertifyMan,
		Status:     in.Status,
		Detail:     in.Detail,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductVertifyRecordAddResp{}, nil
}
