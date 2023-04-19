package module

import (
	"errors"

	"gorm.io/gorm"
	"org.chatgin/src/config"
)

type UserBasic struct {
	gorm.Model
	Account  string //账号
	Username string // 用户明
	Passwrod string // 登录密码
	State    int16  // 用户状态， -1 - 异常；0 - 锁定；1 - 正常；
}

// 根据用户名获取用户信息
func (u *UserBasic) GetByUsername(account string) error {
	if err := config.MysqlDB.Where("account = ? ", account).First(u).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil

}

// 根据ID获取用户信息
func (u *UserBasic) GetById(id uint) error {
	if err := config.MysqlDB.Where("id = ?", id).First(u).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// 新增数据
func (u *UserBasic) Insert() error {
	if err := config.MysqlDB.Create(u).Error; err != nil {
		return err
	}
	return nil
}
