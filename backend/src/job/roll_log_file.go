package job

import (
	"backend/src/constants"
	"backend/src/global"
	"backend/src/utils"
	"log"
	"os"
)

// 使日志滚动起来
func ResetRollLogFile() {
	log.SetPrefix(global.Config.Log.Prefix)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if global.Env == constants.Production {
		logFilePath := utils.GetFullLogPath(global.Config.Log.InfoPath)
		logFile, logErr := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if logErr != nil {
			log.Panic(logErr)
		}
		log.SetOutput(logFile)
	}
	log.Println("init log finished...")
}
