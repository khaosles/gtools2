package result

import "fmt"

/*
   @File: error_code.go
   @Author: khaosles
   @Time: 2023/3/7 21:54
   @Desc:
*/

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e ErrorCode) SetMsg(msg any) ErrorCode {
	e.Msg = fmt.Sprintf("%v", msg)
	return e
}

func (e ErrorCode) Error() string {
	return e.Msg
}

func NewAssertError(msg any) ErrorCode {
	return ErrorCode{Code: 40000}.SetMsg(msg)
}

func NewAuthError(msg any) ErrorCode {
	return ErrorCode{Code: 40100}.SetMsg(msg)
}

func NewNotFoundError(msg any) ErrorCode {
	return ErrorCode{Code: 40400}.SetMsg(msg)
}

func NewForbiddenError(msg any) ErrorCode {
	return ErrorCode{Code: 40300}.SetMsg(msg)
}

func NewInternalError(msg any) ErrorCode {
	return ErrorCode{Code: 50000}.SetMsg(msg)
}

var (
	SUCCESS              = ErrorCode{20000, "ok"}      // 运行成功
	PARAMS_ERROR         = ErrorCode{40000, "请求参数错误"}  // 参数错误
	DUPLICATE_USERNAME   = ErrorCode{40001, "用户名已存在"}  // 参数错误
	NOT_LOGIN_ERROR      = ErrorCode{40100, "未登录"}     // 账号未登录
	NO_AUTH_ERROR        = ErrorCode{40101, "账号无权限"}   // 账号无权限
	TOKEN_EXPIRE         = ErrorCode{40102, "登录信息过期"}  // 登录过期
	TOKEN_REMOTING_LOGIN = ErrorCode{40103, "账号异地登录"}  // 异地登录
	NOT_FOUND_ERROR      = ErrorCode{40400, "请求数据不存在"} // 数据不存在
	FORBIDDEN_ERROR      = ErrorCode{40300, "禁止访问"}    // 数据禁止访问
	LIMIT_ERROR          = ErrorCode{40301, "限制访问"}    // 数据禁止访问
	SYSTEM_ERROR         = ErrorCode{50000, "操作失败"}    // 系统内部出错
)
