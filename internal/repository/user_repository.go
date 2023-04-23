package repository

import (
	"errors"

	"gorm.io/gorm"
	"org.chatgin/internal/config"
)

type UserBasic struct {
	gorm.Model
	Account  string //账号
	Username string // 用户明
	Passwrod string // 登录密码
	State    int16  // 用户状态， -1 - 异常；0 - 锁定；1 - 正常；
}

type UserRepository struct {
}

// 根据用户名获取用户信息
func (u *UserRepository) GetByUsername(account string) (user UserBasic, err error) {
	var userBasci UserBasic
	if err := config.MysqlDB.Where("account = ? ", account).First(userBasci).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return userBasci, err
	}
	return userBasci, nil

}

// 根据ID获取用户信息
func (u *UserRepository) GetById(id uint) (user UserBasic, err error) {
	var userBasci UserBasic
	if err := config.MysqlDB.Where("id = ?", id).First(u).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return userBasci, err
	}
	return userBasci, err
}

// 新增数据
func (u *UserRepository) Insert(user UserBasic) error {
	if err := config.MysqlDB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
