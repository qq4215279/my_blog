package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 在文件末尾追加内容
func Append(fileName, content string) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_SYNC, 0660)
	if nil != err {
		log.Panicf("打开%s失败", fileName)
		return
	}
	f.WriteString(content)
	defer f.Close()
}

// 写入一个文件
func WriteFile(fileName string, content []byte) {
	MkdirAll(GetParentDir(fileName))
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_SYNC, 0660)
	if nil != err {
		log.Panicf("打开%s失败", fileName)
		return
	}
	f.Write(content)
	defer f.Close()
}

// 获取今天日志文件名称
func getToDayLogFileName() string {
	return fmt.Sprintf("%d-%d-%d.log", time.Now().Year(), time.Now().Month(), time.Now().Day())
}

// 创建目录
func MkdirAll(path string) {
	e := os.MkdirAll(path, 0777)
	if e != nil {
		panic(e)
	}
}
