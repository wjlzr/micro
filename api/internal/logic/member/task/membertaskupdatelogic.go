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

type MemberTaskUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskUpdateLogic {
	return MemberTaskUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskUpdateLogic) MemberTaskUpdate(req types.UpdateMemberTaskReq) (*types.UpdateMemberTaskResp, error) {
	_, err := l.svcCtx.Ums.MemberTaskUpdate(l.ctx, &umsclient.MemberTaskUpdateReq{
		Id:           req.Id,
		Name:         req.Name,
		Growth:       req.Growth,
		Intergration: req.Intergration,
		Type:         req.Type,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新会员任务信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新会员任务失败")
	}

	return &types.UpdateMemberTaskResp{
		Code:    "000000",
		Message: "更新会员任务成功",
	}, nil
}
