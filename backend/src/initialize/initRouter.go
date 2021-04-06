package initialize

import (
	"backend/src/constants"
	"backend/src/global"
	"backend/src/handler"
)

func InitGenRouter() {
	apiGroup := global.Engine.Group(constants.BaseUrl)

	// 添加标签
	apiGroup.POST("tag", handler.AddTag)
	// 获取所有标签
	apiGroup.GET("tag", handler.GetAllTag)
	// 修改标签
	apiGroup.PUT("tag", handler.UpdateTag)
	// 删除标签
	apiGroup.DELETE("tag", handler.DeleteTag)

	// 添加分类
	apiGroup.POST("category", handler.AddCategory)
	// 获取所有分类
	apiGroup.GET("category", handler.GetAllCategory)
	// 修改分类
	apiGroup.PUT("category", handler.UpdateCategory)
	// 删除分类
	apiGroup.DELETE("category", handler.DeleteCategory)

	// 添加博客
	apiGroup.POST("blog", handler.AddBlog)
	// 获取博客
	apiGroup.GET("blog", handler.GetBlogById)
	// 修改博客
	apiGroup.PUT("blog", handler.UpdateBlog)
	// 删除博客
	apiGroup.DELETE("blog", handler.DeleteBlog)

	// 登录
	apiGroup.POST("login", handler.Login)
	// 登出
	apiGroup.POST("logout", handler.Logout)
	// 刷新accessToken
	apiGroup.POST("refreshAccessToken", handler.RefreshAccessToken)
	// token登录
	apiGroup.POST("loginByToken", handler.LoginByToken)
	// 禁用用户
	apiGroup.POST("disableUser", handler.DisableUser)
	// 启用用户
	apiGroup.Group("user").POST("enableUser", handler.EnableUser)
	// 修改ldap登陆方式的admin账号的密码
	apiGroup.GET("updateLdapAdminPassword", handler.UpdateLdapAdminPassword)

	// 上传文件
	apiGroup.POST("uploadFile", handler.UploadFile)
	// 下载文件
	apiGroup.GET("downloadFile", handler.DownloadFile)
	// 图片上传
	apiGroup.POST("uploadPicture", handler.UploadPicture)
	// 图片下载
	apiGroup.GET("downloadPicture", handler.FindPicture)

}
