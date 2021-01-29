/**
 * blog_service
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/29 17:46
 */
package service

import (
	"backend/src/module"
	"backend/src/repository"
	"backend/src/utils"
	"reflect"
)

// 保存博客
func AddBlog(param module.Blog) ResponseBody {
	if utils.IsBlank(param.Title) {
		return NewCustomErrorResponseBody("标题不能为空")
	}
	if utils.IsBlank(param.Summary) {
		return NewCustomErrorResponseBody("摘要不能为空")
	}
	if param.CategoryId < 1 || utils.IsBlank(param.TagIds) {
		return NewCustomErrorResponseBody("参数不完整")
	}

	repository.Create(&param)

	return NewSimpleSuccessResponseBody()
}

// 获取博客
func GetBlogById(blogId uint) ResponseBody {
	var blog module.Blog
	repository.FindById(&blog, blogId)

	if reflect.DeepEqual(blog, module.Blog{}) {
		return NewParamErrorResponseBody()
	}

	return NewSuccessResponseBody(blog)
}

// 修改博客
func UpdateBlog(param module.Blog) ResponseBody {
	var blog module.Blog
	repository.FindById(&blog, param.Id)

	if reflect.DeepEqual(blog, module.Blog{}) {
		return NewParamErrorResponseBody()
	}

	if !utils.IsBlank(param.Title) {
		blog.Title = param.Title
	}
	if !utils.IsBlank(param.Summary) {
		blog.Summary = param.Summary
	}
	if !utils.IsBlank(param.Contend) {
		blog.Contend = param.Contend
	}
	if param.CategoryId > 0 {
		blog.CategoryId = param.CategoryId
	}
	if !utils.IsBlank(param.TagIds) {
		blog.TagIds = param.TagIds
	}
	if !utils.IsBlank(param.ImageUrl) {
		blog.ImageUrl = param.ImageUrl
	}

	repository.Save(&blog)

	return NewSimpleSuccessResponseBody()
}

// 删除博客
func DeleteBlog(blogId uint) ResponseBody {
	var blog module.Blog
	repository.FindById(&blog, blogId)

	if reflect.DeepEqual(blog, module.Blog{}) {
		return NewParamErrorResponseBody()
	}

	repository.Delete(&blog)

	return NewSimpleSuccessResponseBody()
}
