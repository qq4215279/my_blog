package main

import (
	"backend/src/core"
	"backend/src/initialize"
)

// 初始化
func init() {
	initialize.Init()
}

// 主函数
func main() {
	core.RunServer()
}
