package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hario/ci/server/constants"
	"hario/ci/server/global"
	"time"
)

// 根据环境启动服务
func RunServer() {
	if global.Env == constants.Production {
		RunUnixServer()
	} else {
		RunWindowsServer()
	}
}

// 启动服务
func runServer() {
	global.Engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	_ = global.Engine.Run(global.Config.System.Addr)
}
