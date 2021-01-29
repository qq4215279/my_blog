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

}
