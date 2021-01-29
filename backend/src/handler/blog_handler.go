/**
 * blog_handler
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/29 17:43
 */
package handler

import (
	"backend/src/module"
	"backend/src/service"
	"github.com/gin-gonic/gin"
)

// 保存博客
func AddBlog(ctx *gin.Context) {
	var param module.Blog
	_ = ctx.ShouldBind(&param)

	WrapperResponseBody(ctx, service.AddBlog(param))
}

// 获取博客
func GetBlogById(ctx *gin.Context) {
	var blogId = PostFormUIntValue(ctx, "id")

	WrapperResponseBody(ctx, service.GetBlogById(blogId))
}

// 修改博客
func UpdateBlog(ctx *gin.Context) {
	var param module.Blog
	_ = ctx.ShouldBind(&param)

	WrapperResponseBody(ctx, service.UpdateBlog(param))
}

// 删除博客
func DeleteBlog(ctx *gin.Context) {
	var blogId = PostFormUIntValue(ctx, "id")

	WrapperResponseBody(ctx, service.DeleteBlog(blogId))
}
