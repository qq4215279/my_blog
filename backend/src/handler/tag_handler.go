/**
 * tag_handler
 * @author liuzhen
 * @Description 标签handle
 * @version 1.0.0 2021/1/28 17:04
 */
package handler

import (
	"backend/src/module"
	"backend/src/service"
	"github.com/gin-gonic/gin"
)

// 添加标签
func AddTag(ctx *gin.Context) {
	var param module.Tag
	_ = ctx.ShouldBind(&param)

	result := service.AddTag(param)
	WrapperResponseBody(ctx, result)
}

//获取所有标签
func GetAllTag(ctx *gin.Context) {
	WrapperResponseBody(ctx, service.GetAllTag())
}

// 修改标签
func UpdateTag(ctx *gin.Context) {
	var param module.Tag
	ctx.ShouldBind(&param)
	WrapperResponseBody(ctx, service.UpdateTag(param))
}

// 删除标签
func DeleteTag(ctx *gin.Context) {
	var tagId = PostFormUIntValue(ctx, "id")
	WrapperResponseBody(ctx, service.DeleteTag(tagId))
}
