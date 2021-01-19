package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 成功响应
type SuccessResponse struct {
	State uint        `json:"state"`
	Data  interface{} `json:"data"`
}

// 错误响应
type ErrorResponse struct {
	State uint   `json:"state"`
	Msg   string `json:"msg"`
}

// 获取数字类型参数
func PostFormIntValue(context *gin.Context, key string) (value int) {
	intValue, _ := strconv.Atoi(context.PostForm(key))
	value = intValue
	return value
}

// 获取数字类型参数
func PostFormUIntValue(context *gin.Context, key string) uint {
	return uint(PostFormIntValue(context, key))
}

func WrapperResponseBody(context *gin.Context, body service.ResponseBody) {
	if body.State != constants.ResponseStateSuccess {
		context.JSON(http.StatusOK, gin.H{
			"state": body.State,
			"msg":   body.Msg,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"state": body.State,
			"data":  body.Data,
		})
	}
}
