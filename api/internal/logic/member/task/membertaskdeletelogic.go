package logic

import (
	"context"
	"micro/common/errorx"
	"micro/rpc/svs/ums/umsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskDeleteLogic {
	return MemberTaskDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskDeleteLogic) MemberTaskDelete(req types.DeleteMemberTaskReq) (*types.DeleteMemberTaskResp, error) {
	_, err := l.svcCtx.Ums.MemberTaskDelete(l.ctx, &umsclient.MemberTaskDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除会员任务异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除会员任务失败")
	}

	return &types.DeleteMemberTaskResp{
		Code:    "000000",
		Message: "",
	}, nil
}
