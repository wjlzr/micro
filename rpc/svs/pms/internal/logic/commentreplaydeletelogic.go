package logic

import (
	"context"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentReplayDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentReplayDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentReplayDeleteLogic {
	return &CommentReplayDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentReplayDeleteLogic) CommentReplayDelete(in *pms.CommentReplayDeleteReq) (*pms.CommentReplayDeleteResp, error) {
	err := l.svcCtx.PmsCommentReplayModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.CommentReplayDeleteResp{}, nil
}
