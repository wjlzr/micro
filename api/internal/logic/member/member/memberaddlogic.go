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

type MemberAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddLogic {
	return MemberAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddLogic) MemberAdd(req types.AddMemberReq) (*types.AddMemberResp, error) {
	_, err := l.svcCtx.Ums.MemberAdd(l.ctx, &umsclient.MemberAddReq{
		MemberLevelId:         req.MemberLevelId,
		Username:              req.Username,
		Password:              req.Password,
		Nickname:              req.Nickname,
		Phone:                 req.Phone,
		Status:                req.Status,
		CreateTime:            req.CreateTime,
		Icon:                  req.Icon,
		Gender:                req.Gender,
		Birthday:              req.Birthday,
		City:                  req.City,
		Job:                   req.Job,
		PersonalizedSignature: req.PersonalizedSignature,
		SourceType:            req.SourceType,
		Integration:           req.Integration,
		Growth:                req.Growth,
		LuckeyCount:           req.LuckeyCount,
		HistoryIntegration:    req.HistoryIntegration,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员信息失败")
	}

	return &types.AddMemberResp{
		Code:    "000000",
		Message: "添加会员信息成功",
	}, nil
}
