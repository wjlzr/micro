package logic

import (
	"context"

	"micro/rpc/svs/sys/internal/svc"
	"micro/rpc/svs/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type JobDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobDeleteLogic {
	return &JobDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JobDeleteLogic) JobDelete(in *sys.JobDeleteReq) (*sys.JobDeleteResp, error) {
	err := l.svcCtx.JobModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sys.JobDeleteResp{}, nil
}
