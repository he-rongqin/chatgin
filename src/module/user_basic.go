package module

import (
	"gorm.io/gorm"
	"org.chatgin/src/config"
)

type UserBasic struct {
	gorm.Model
	account  string //账号
	Username string // 用户明
	Passwrod string // 登录密码
	State    int16  // 用户状态， -1 - 异常；0 - 锁定；1 - 正常；
}

// 创建表
func (user *UserBasic) CreateTable() {
	config.MysqlDB.AutoMigrate(&UserBasic{})
}

// 根据用户名获取用户信息
func (u *UserBasic) GetByUsername(account string) *UserBasic {
	config.MysqlDB.Where("account = ? ", account).First(u)
	return u
}

// 新增数据
func (u *UserBasic) Insert(user UserBasic) bool {
	config.MysqlDB.Create(user)
	return true
}
