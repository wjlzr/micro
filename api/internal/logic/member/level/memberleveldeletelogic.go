package logic

import (
	"context"
	"micro/common/errorx"
	"micro/rpc/svs/ums/umsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLevelDeleteLogic {
	return MemberLevelDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLevelDeleteLogic) MemberLevelDelete(req types.DeleteMemberLevelReq) (*types.DeleteMemberLevelResp, error) {
	_, err := l.svcCtx.Ums.MemberLevelDelete(l.ctx, &umsclient.MemberLevelDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除会员等级异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除会员等级失败")
	}
	return &types.DeleteMemberLevelResp{
		Code:    "000000",
		Message: "添加会员等级成功",
	}, nil
}
