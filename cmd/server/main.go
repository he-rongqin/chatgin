package main

import (
	"fmt"

	"org.chatgin/internal/config"
	"org.chatgin/internal/router"
	"org.chatgin/pkg/database"
)

// 初始化方法
func init() {
	// 初始化配置
	config.InitConfig()
	// 初始化表结构
	database.InitTables()
}

func main() {

	r := router.Router()
	r.Run(config.ServerConfig.ServerPort)
	fmt.Printf("%v 启动完成。\n", config.ServerConfig.AppName)
}
