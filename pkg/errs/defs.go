package errs

import "net/http"

type HttpErrCode = int32

const (
	App     HttpErrCode = 10001
	Device  HttpErrCode = 10002
	Group   HttpErrCode = 10003
	Message HttpErrCode = 10004
	User    HttpErrCode = 10005
)

type Err struct {
	Code   HttpErrCode `json:"error_code"`
	ErrMsg string      `json:"msg"`
}

type HttpErr struct {
	HttpSC int
	Err
}

func (err Err) Error() string {
	return err.ErrMsg
}

var (
	ServerInternalError = HttpErr{HttpSC: http.StatusInternalServerError, Err: Err{Code: 0, ErrMsg: "Internal service error"}}
	ParameterError      = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10000, ErrMsg: "参数验证失败！"}}
)

func NewHttpErr(code HttpErrCode, errMsg string) HttpErr {
	return HttpErr{
		HttpSC: http.StatusOK,
		Err: Err{
			Code:   code,
			ErrMsg: errMsg,
		},
	}
}
