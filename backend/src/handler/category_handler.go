/**
 * category_handler
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/29 16:58
 */
package handler

import (
	"backend/src/module"
	"backend/src/service"
	"github.com/gin-gonic/gin"
)

// 添加分类
func AddCategory(ctx *gin.Context) {
	var param module.Category
	_ = ctx.ShouldBind(&param)

	WrapperResponseBody(ctx, service.AddCategory(param))
}

// 获取所有分类
func GetAllCategory(ctx *gin.Context) {
	WrapperResponseBody(ctx, service.GetAllCategory())
}

// 修改分类信息
func UpdateCategory(ctx *gin.Context) {
	var param module.Category
	_ = ctx.ShouldBind(&param)

	WrapperResponseBody(ctx, service.UpdateCategory(param))
}

// 删除分类
func DeleteCategory(ctx *gin.Context) {
	categoryId := PostFormUIntValue(ctx, "id")

	WrapperResponseBody(ctx, service.DeleteCategory(categoryId))
}
