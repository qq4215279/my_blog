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

}
