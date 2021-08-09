package logic

import (
	"context"

	"micro/rpc/svs/ums/internal/svc"
	"micro/rpc/svs/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingDeleteLogic {
	return &MemberRuleSettingDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingDeleteLogic) MemberRuleSettingDelete(in *ums.MemberRuleSettingDeleteReq) (*ums.MemberRuleSettingDeleteResp, error) {
	err := l.svcCtx.UmsMemberRuleSettingModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberRuleSettingDeleteResp{}, nil
}
