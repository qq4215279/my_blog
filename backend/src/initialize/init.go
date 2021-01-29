package initialize

import (
	"backend/src/cache"
	"backend/src/config"
	"backend/src/constants"
	"backend/src/global"
	"backend/src/job"
	"backend/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	runtimeEnv = ""
)

// 初始化
func Init() {
	initSystem()
	initGlobal()
	initTask()
	initJob()
}

// 初始化全局变量
func initGlobal() {
	global.Engine = initRouters()
	InitGenRouter()
	log.Println("init gen router finished...")
	global.DataBase = initGorm().Debug()
	log.Println("init global finished...")
}

func initTask() {
	log.Printf("init task")
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for range ticker.C {
			cache.ClearUser()
		}
	}()
}

// 初始化任务
func initJob() {
	global.Cron = cron.New()
	global.Cron.AddFunc(global.Config.Job.Clear, job.Clear)
	global.Cron.AddFunc("0 0 0 * * *", job.ResetRollLogFile)
	global.Cron.Start()
}

// 初始化系统
func initSystem() {
	initPath()
	initConfig()
	initDir()
	initLog()
	rand.Seed(time.Now().UnixNano())
}

// 初始化路径
func initPath() {
	workingDirectory := utils.GetParentDir(utils.GetWd())
	os.Chdir(workingDirectory)
	log.Println("init path finished, Current Working Directory: ", workingDirectory)
}

// 初始化配置
func initConfig() {
	yamlFile, readFileErr := ioutil.ReadFile(filepath.Join(config.ConfPath, getConfigFileName()))
	if readFileErr != nil {
		panic(readFileErr)
	}
	var conf config.Config
	var parseYamlErr = yaml.Unmarshal(yamlFile, &conf)
	if parseYamlErr != nil {
		panic(parseYamlErr)
	}
	global.Config = conf
	log.Println("init config finished...")
}

// 初始化目录
func initDir() {
	// 不存在就创建目录
	utils.MkdirAll(filepath.Join(utils.GetWd(), global.Config.Log.LogBasePath, global.Config.Log.InfoPath))
	utils.MkdirAll(filepath.Join(utils.GetWd(), global.Config.Log.LogBasePath, global.Config.Log.GinPath))
	utils.MkdirAll(filepath.Join(utils.GetWd(), global.Config.Log.LogBasePath, global.Config.Log.OrmPath))
	utils.MkdirAll(filepath.Join(utils.GetWd(), global.Config.Log.LogBasePath, global.Config.Log.BuildJobLogPath))
	utils.MkdirAll(filepath.Join(utils.GetWd(), global.Config.Build.FilePath))
	utils.MkdirAll(filepath.Join(utils.GetWd(), constants.SSHKeyDir))
	log.Println("init dir finished...")
}

// 初始化日志
func initLog() {
	job.ResetRollLogFile()
	log.Println("init log finished...")
}

// 获取配置文件名
func getConfigFileName() string {
	environ := os.Environ()
	env := ""
	for _, envProp := range environ {
		if !strings.Contains(envProp, "=") {
			continue
		}
		envKey := strings.Split(envProp, "=")[0]
		envValue := strings.Split(envProp, "=")[1]
		if envKey == "CI_ENV" {
			env = envValue
			break
		}
	}
	if utils.IsBlank(env) {
		env = "dev"
	}
	log.Printf("load config %s", "conf_"+env+".yml")
	runtimeEnv = env
	global.Env = env
	return "conf_" + env + ".yml"
}

// 使用gorm连接数据库
func initGorm() *gorm.DB {
	conf := global.Config.Mysql
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", conf.Username, conf.Password, conf.URL, conf.DbName, conf.Config)
	log.Printf("connect db dataSourceName %s", dataSourceName)
	var config gorm.Config

	db, err := gorm.Open(mysql.Open(dataSourceName), &config)
	if err != nil {
		log.Printf("init gorm error %s", err)
		return nil
	}
	log.Println("init gorm finished...")

	return db
}

// 初始化gin路由
func initRouters() *gin.Engine {
	var engine = gin.Default()
	// 跨域配置
	engine.Use(Cors())
	// 日志拦截器
	engine.Use(LogMiddleware())
	// 参数拦截器
	engine.Use(ParamMiddleware())
	// 全局异常捕获
	engine.Use(RecoverMiddleware())
	// 用户拦截器
	//engine.Use(AuthMiddleware())
	// 权限拦截器
	//engine.Use(PrivilegeMiddleware())

	ApiGroup := engine.Group(constants.BaseUrl)
	ApiGroup.Use(ApiMiddlewareTest())
	log.Println("init gin router finished...")
	return engine
}

type ORMWriter struct {
}

func (ow *ORMWriter) Printf(format string, args ...interface{}) {
	// TODO io操作过于频繁
	logFile := utils.GetFullLogPath(global.Config.Log.OrmPath)
	utils.WriteFile(logFile, []byte(fmt.Sprintf(format, args...)))
}
