package service

import "constants"

type ResponseBody struct {
	Data  interface{} `json:"data"`
	State int         `json:"state"`
	Msg   string      `json:"msg"`
}

// 创建自定义异常
func NewCustomErrorResponseBody(msg string) ResponseBody {
	return newError(msg)
}

// 创建成功返回
func NewSuccessResponseBody(data interface{}) ResponseBody {
	return ResponseBody{State: constants.ResponseStateSuccess, Data: data}
}

// 创建简单操作成功返回
func NewSimpleSuccessResponseBody() ResponseBody {
	return ResponseBody{State: constants.ResponseStateSuccess, Data: "操作成功"}
}

// 创建空参数异常返回
func NewParamEmptyResponseBody() ResponseBody {
	return newError("参数不能为空")
}

// 创建参数错误异常返回
func NewParamErrorResponseBody() ResponseBody {
	return newError("参数错误")
}

// 创建重新登陆异常返回
func NewNonSessionErrorResponseBody() ResponseBody {
	return ResponseBody{State: constants.ResponseStateRedirect, Msg: "请重新登录"}
}

// 创建没有权限异常返回
func NewNonPrivilegeErrorResponseBody() ResponseBody {
	return newError("没有权限")
}

// 创建没有权限异常返回
func NewNonFuncPrivilegeErrorResponseBody() ResponseBody {
	return newError("没有功能权限")
}

// 创建没有权限异常返回
func NewNonModulePrivilegeErrorResponseBody() ResponseBody {
	return newError("没有模块权限")
}

// 创建异常返回
func newError(msg string) ResponseBody {
	return ResponseBody{State: constants.ResponseStateFail, Msg: msg}
}
