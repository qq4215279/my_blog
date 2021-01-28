package global

import (
	"backend/src/config"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

// 全局变量
var (
	// 配置文件
	Config config.Config
	// gin引擎
	Engine *gin.Engine
	// 数据库
	DataBase *gorm.DB
	// 环境
	Env string
	// 环境
	Cron *cron.Cron
)
