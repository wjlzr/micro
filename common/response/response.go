package response

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"micro/common/errorx"
	"net/http"
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 错误返回的json
func Error(code int) error {
	return NewCodeError(code)
}

func NewCodeError(code int) error {
	return &CodeError{Code: code, Msg: errorx.GetMsg(code)}
}

//正确返回的json
func Success(w http.ResponseWriter, v interface{}) {
	httpx.OkJson(w, Response{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: v,
	})
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

func (e *CodeError) Error() string {
	return e.Msg
}
