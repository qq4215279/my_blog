package job

import (
	"backend/src/global"
	"backend/src/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// 清理任务
func Clear() {
	log.Printf("Clear job start")
	// 清理日志
	e := clearFile(utils.GetLogRootPath(), global.Config.Job.ClearLogRetentionDay)
	if e != nil {
		log.Printf("clear log error %s", e.Error())
	}
	// 清理打包结果
	e = clearFile(utils.GetBuildFileRootPath(), global.Config.Job.ClearBuildFileRetentionDay)
	if e != nil {
		log.Printf("clear build file error %s", e.Error())
	}
	e = clearEmptyDir(utils.GetBuildFileRootPath())
	if e != nil {
		log.Printf("clear build dir error %s", e.Error())
	}
	log.Printf("Clear job success")
}

// 清理文件
func clearFile(path string, days int) error {
	dir, e := ioutil.ReadDir(path)
	if e != nil {
		return e
	}
	for _, f := range dir {
		fullPath := filepath.Join(path, f.Name())
		if f.IsDir() {
			e := clearFile(fullPath, days)
			if e != nil {
				log.Printf("clear file %s faile %s", fullPath, e.Error())
			}
		} else {
			subDays := int(time.Now().Sub(f.ModTime()).Hours()) / 24
			log.Printf("name %s modfiy time %s sub days %d", f.Name(), f.ModTime(), subDays)
			if subDays >= days {
				e = os.Remove(fullPath)
				if e != nil {
					log.Printf("remove file %s faile %s", fullPath, e.Error())
				}
				log.Printf("remove file %s", fullPath)
			}
		}
	}
	return e
}

// 清理空目录
func clearEmptyDir(path string) error {
	dir, e := ioutil.ReadDir(path)
	if e != nil {
		return e
	}
	for _, f := range dir {
		fullPath := filepath.Join(path, f.Name())
		if f.IsDir() && f.Size() == 0 {
			dir, e := ioutil.ReadDir(fullPath)
			if len(dir) == 0 {
				e = os.Remove(fullPath)
				if e != nil {
					log.Printf("remove emptyDir %s faile %s", fullPath, e.Error())
				}
				log.Printf("remove emptyDir %s", fullPath)
			}
		} else {
			_ = os.Remove(fullPath)
		}
	}
	return e
}
