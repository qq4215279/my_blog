package utils

import (
	"backend/src/global"
	"os"
	"path/filepath"
)

// 获取文件名
func GetFileName(filePath string) string {
	_, file := filepath.Split(filePath)
	return file
}

func GetWd() string {
	wd, _ := os.Getwd()
	return wd
}

// 获取父目录
func GetParentDir(filePath string) string {
	parentDir, _ := filepath.Split(filePath)
	return parentDir
}

// 获取日志文件全路径
func GetFullLogPath(logPath string) string {
	return filepath.Join(GetWd(), global.Config.Log.LogBasePath, logPath, getToDayLogFileName())
}

// 获取日志文件根目录
func GetLogRootPath() string {
	return filepath.Join(GetWd(), global.Config.Log.LogBasePath)
}

// 获取打包结果根目录
func GetBuildFileRootPath() string {
	return filepath.Join(GetWd(), global.Config.Build.FilePath)
}

// 获取文件上传目录
func GetUploadFilePath() string {
	return filepath.Join(GetWd(), global.Config.Upload.UploadBasePath, global.Config.Upload.FilePath)
}

// 获取图片上传目录
func GetUploadPicturePath() string {
	return filepath.Join(GetWd(), global.Config.Upload.UploadBasePath, global.Config.Upload.PicturePath)
}
