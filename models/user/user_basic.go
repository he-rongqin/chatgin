package user

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Username string // 用户明
	Passwrod string // 登录密码
	State    int16  // 用户状态， -1 - 异常；0 - 锁定；1 - 正常；
}
