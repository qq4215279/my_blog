package initialize

import (
	"backend/src/cache"
	"backend/src/constants"
	"backend/src/handler"
	"backend/src/service"
	"backend/src/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

var (
	// 不用校验用户的白名单前缀
	whiteListPrefix = []string{
		"user/login",
		"user/loginByToken",
		"user/getSubAccountList",
		"system/",
		"callBack/",
		"client/",
		"plugin/",
		"ws/",
		"buildRecord/readLog",
		"buildFlow/getList",
		"env/getEnvSelect",
		"key/downloadSSHKey",
	}
	// 不需要打印接口参数日志的白名单前缀
	whiteListNoLogPrefix = []string{"buildRecord/appendLog"}
)

// 跨域配置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Accept,Authorization, Content-Type, Content-Length, Origin, X-Requested-With, X-CSRF-Token, Token, session")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Cache-Control, Content-Language, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			// 设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

// 日志中间件
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !needPrintMiddlewareLog(c.FullPath()) {
			c.Next()
			return
		}
		start := time.Now().UnixNano() / 1e6
		c.Next()
		end := time.Now().UnixNano() / 1e6
		// url#paramsMap
		log.Printf("%s#%s#%dms", c.Request.URL, c.Request.Form, end-start)
	}
}

// 参数处理中间件
func ParamMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountInfo, _ := c.Cookie("op_account_info")
		var account = utils.String2Map(accountInfo)
		sessionId, _ := c.Cookie("session_id")
		c.Set(constants.SessionIdKey, sessionId)
		c.Set(constants.UserNameKey, account["userName"])
		c.Next()
	}
}

// recover中间件
func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("stack %s", string(debug.Stack()))
				log.Printf("panic:%v\n", r)
				debug.PrintStack()
				handler.WrapperResponseBody(c, service.NewCustomErrorResponseBody("recover success"))
				// 终止调用
				log.Printf("%s call %s fail %s", c.Request.URL, c.FullPath(), "recover")
				c.Abort()
			}
		}()
		c.Next()
	}
}

// apiMiddlewareTest
func ApiMiddlewareTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// 用户拦截中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不需要校验就放行
		if !needCheckAuth(c.FullPath()) {
			c.Next()
			return
		}
		sessionId, _ := c.Cookie("session_id")
		_, ok := cache.GetOnlineUser()[sessionId]
		if ok {
			cache.AddInvalidTime(sessionId)
			c.Next()
			return
		}
		log.Printf("%s call %s fail %s", c.Request.URL, c.FullPath(), "need login")
		c.Abort()
		handler.WrapperResponseBody(c, service.NewNonSessionErrorResponseBody())
	}
}

// 权限拦截中间件
func PrivilegeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不需要校验就放行
		if !needCheckAuth(c.FullPath()) {
			c.Next()
			return
		}
		sessionId, _ := c.Cookie("session_id")
		_, ok := cache.GetOnlineUser()[sessionId]
		if !ok {
			log.Printf("%s call %s fail %s", c.Request.URL, c.FullPath(), "need login")
			c.Abort()
			handler.WrapperResponseBody(c, service.NewNonSessionErrorResponseBody())
			return
		}
		/*u, hasUser := cache.GetUser(sessionId)
		if !hasUser {
			log.Printf("%s call %s fail %s", c.Request.URL, c.FullPath(), "need login")
			c.Abort()
			handler.WrapperResponseBody(c, service.NewNonSessionErrorResponseBody())
			return
		}*/
		// TODO 这个url是不是属于某个function
		/*if service.NeedCheckApiPrivilege(c.FullPath()) {
			// 有没有这个function的权限
			if !service.CanAccessApi(u.UserName, c.FullPath()) {
				log.Printf("%s call %s fail %s", c.Request.URL, c.FullPath(), "need privilege")
				c.Abort()
				handler.WrapperResponseBody(c, service.NewNonFuncPrivilegeErrorResponseBody())
				return
			}
			c.Next()
			return
		}*/
		// TODO 校验模块权限
		/*if !service.HasModulePrivilege(u.UserName, c.FullPath()) {
			log.Printf("%s call %s fail %s", c.Request.URL, c.FullPath(), "need module privilege")
			c.Abort()
			handler.WrapperResponseBody(c, service.NewNonModulePrivilegeErrorResponseBody())
			return
		}*/
		c.Next()
		return
	}
}

// 是否要校验用户session
func needCheckAuth(fullPath string) bool {
	for _, whiteUrlPrefix := range whiteListPrefix {
		if strings.HasPrefix(fullPath, constants.BaseUrl+whiteUrlPrefix) {
			return false
		}
	}
	return true
}

// 是否要校验用户session
func needPrintMiddlewareLog(fullPath string) bool {
	for _, whiteUrlPrefix := range whiteListNoLogPrefix {
		if strings.HasPrefix(fullPath, constants.BaseUrl+whiteUrlPrefix) {
			return false
		}
	}
	return true
}
