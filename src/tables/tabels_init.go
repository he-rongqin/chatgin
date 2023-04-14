package tables

import (
	"org.chatgin/src/config"
	"org.chatgin/src/module"
)

// 初始化表结构，如果存在则不创建
func InitTables() {
	config.MysqlDB.AutoMigrate(&module.UserBasic{})

}
