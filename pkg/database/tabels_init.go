package database

import (
	"org.chatgin/internal/config"
	"org.chatgin/internal/repository"
)

// 初始化表结构，如果存在则不创建
func InitTables() {
	config.MysqlDB.AutoMigrate(&repository.UserBasic{})

}
