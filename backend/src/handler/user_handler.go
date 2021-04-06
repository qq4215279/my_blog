package handler

import (
	"backend/src/constants"
	"backend/src/module"
	"backend/src/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 登录
func Login(context *gin.Context) {
	var param module.UserVo
	_ = context.ShouldBindWith(&param, binding.Form)
	result := service.Login(param)
	WrapperResponseBody(context, result)
}

// 刷新accessToken
func RefreshAccessToken(context *gin.Context) {
	refreshToken := context.Request.Header.Get("refreshToken")
	userName := context.MustGet("userName").(string)

	result := service.RefreshAccessToken(refreshToken, userName)
	WrapperResponseBody(context, result)
}

// token登录
func LoginByToken(context *gin.Context) {
	var param module.UserVo
	_ = context.ShouldBindWith(&param, binding.Form)

	// ldap登录
	sessionId := context.MustGet(constants.SessionIdKey)
	result := service.LoginByLADPToken(param.Token, sessionId.(string))
	WrapperResponseBody(context, result)
}

// 登出
func Logout(context *gin.Context) {
	userName := context.MustGet(constants.UserNameKey).(string)
	sessionId := context.MustGet(constants.SessionIdKey).(string)
	result := service.Logout(userName, sessionId)
	WrapperResponseBody(context, result)
}

// 禁用用户
func DisableUser(context *gin.Context) {
	userName := context.PostForm("disableUserName")
	result := service.DisableUser(userName)
	WrapperResponseBody(context, result)
}

// 启用用户
func EnableUser(context *gin.Context) {
	userName := context.PostForm("enableUserName")
	result := service.EnableUser(userName)
	WrapperResponseBody(context, result)
}

// 修改ldap登陆方式的admin账号的密码
func UpdateLdapAdminPassword(context *gin.Context) {
	var param module.UserVo
	_ = context.ShouldBindWith(&param, binding.Form)
	result := service.UpdateLdapAdminPassword(param)
	WrapperResponseBody(context, result)
}
