package logic

import (
	"context"
	"micro/rpc/model/pmsmodel"
	"time"

	"micro/rpc/svs/pms/internal/svc"
	"micro/rpc/svs/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentReplayAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentReplayAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentReplayAddLogic {
	return &CommentReplayAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentReplayAddLogic) CommentReplayAdd(in *pms.CommentReplayAddReq) (*pms.CommentReplayAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.PmsCommentReplayModel.Insert(pmsmodel.PmsCommentReplay{
		CommentId:      in.CommentId,
		MemberNickName: in.MemberNickName,
		MemberIcon:     in.MemberIcon,
		Content:        in.Content,
		CreateTime:     CreateTime,
		Type:           in.Type,
	})
	if err != nil {
		return nil, err
	}

	return &pms.CommentReplayAddResp{}, nil
}
