package user

import (
	"context"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	"micro/api/internal/svc"
	"micro/api/internal/types"
	"micro/common/errorx"
	"micro/common/response"
	"micro/rpc/svs/oauth/oauthclient"
	"strings"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 根据用户名和密码登录
func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {

	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("用户名或密码不能为空,请求信息失败,参数:%s", reqStr)
		return nil, response.Error(errorx.ERROR_USERNAME_PASSWORD_NOT_EMPTY)
	}

	resp, err := l.svcCtx.Oauth.Login(l.ctx, &oauthclient.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据用户名: %s和密码: %s查询用户异常:%s", req.UserName, req.Password, err.Error())
		return nil, response.Error(errorx.ERROR_FIND_USER_EXEPTION)
	}

	return &types.LoginResp{
		CurrentAuthority: resp.CurrentAuthority,
		Id:               resp.Id,
		UserName:         resp.UserName,
		AccessToken:      resp.AccessToken,
		AccessExpire:     resp.AccessExpire,
		RefreshAfter:     resp.RefreshAfter,
	}, nil
}
