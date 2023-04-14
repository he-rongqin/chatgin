package service

import (
	"errors"

	"org.chatgin/src/modules/module"
)

type UserService struct {
}

// 用户登录表单
type UserLoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Paasword string `form:"paasword" json:"password" binding:"required"`
}

type UserInfo struct {
	id       uint
	username string
	state    int16
}

func NewUserInfo(id uint, state int16, username string) *UserInfo {
	return &UserInfo{
		id:       id,
		state:    state,
		username: username,
	}
}

// user login
func (u *UserService) Login(userLogin UserLoginForm) (userInfo *UserInfo, erro error) {
	if userLogin.Username == "" {
		return nil, errors.New("登录名不允许为空")
	}
	if userLogin.Paasword == "" {
		return nil, errors.New("登录密码不允许为空")
	}
	user := &module.UserBasic{}
	user.getByUsername(userLogin.Paasword)
	if user.Username == "" {
		return nil, errors.New("用户名或密码错误")
	}
	// todo 判断密码

	// todo 生成token

	// 返回登录信息
	return NewUserInfo(user.ID, user.State, user.Username), nil
}

// user logout
func (u *UserService) Logout(username string) bool {
	return true

}

// get userinfo
func (u *UserService) GetInfo(id int) *UserInfo {
	return nil
}
