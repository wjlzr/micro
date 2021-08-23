package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest/httpx"
	"micro/common/errorx"
	"net/http"
	"strings"
)

type CheckUrlMiddleware struct {
	Redis *redis.Redis
}

func NewCheckUrlMiddleware(Redis *redis.Redis) *CheckUrlMiddleware {
	return &CheckUrlMiddleware{Redis: Redis}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//判断请求header中是否携带了x-user-id
		userId := r.Context().Value("userId").(json.Number).String()
		if userId == "" {
			logx.Errorf(errorx.GetMsg(errorx.ERROR_NOT_EXIST_PARAME_X_USER_ID))
			httpx.Error(w, errors.New(errorx.GetMsg(errorx.ERROR_NOT_EXIST_PARAME_X_USER_ID)))
			return
		}

		//获取用户能访问的url
		urls, err := m.Redis.Get(userId)
		if err != nil {
			logx.Errorf(errorx.GetMsg(errorx.ERROR_REDIS_CONNECT_EXEPTION), userId)
			httpx.Error(w, errors.New(fmt.Sprintf(errorx.GetMsg(errorx.ERROR_REDIS_CONNECT_EXEPTION), userId)))
			return
		}

		if len(strings.TrimSpace(urls)) == 0 {
			logx.Errorf(errorx.GetMsg(errorx.ERROR_NOT_LOGIN), userId)
			httpx.Error(w, errors.New(errorx.GetMsg(errorx.ERROR_NOT_LOGIN)))
			return
		}

		backUrls := strings.Split(urls, ",")

		b := false
		for _, url := range backUrls {
			if url == r.RequestURI {
				b = true
				break
			}
		}

		if b || r.RequestURI == "/api/sys/user/currentUser" || r.RequestURI == "/api/sys/user/selectAllData" || r.RequestURI == "/api/sys/role/queryMenuByRoleId" {
			logx.Infof("用户userId: %s,访问: %s路径", userId, r.RequestURI)
			next(w, r)
		} else {
			logx.Errorf("用户userId: %s,没有访问: %s路径的权限", userId, r.RequestURI)
			httpx.Error(w, errors.New(errorx.GetMsg(errorx.ERROR_NOT_ACCESS_AUTHORITY)))
			return
		}

	}
}
