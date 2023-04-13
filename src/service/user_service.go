package service

import (
	"errors"
	"net/http"

	"org.chatgin/src/common"
	"org.chatgin/src/modules/module"
)

type UserService struct {
}

// 用户登录表单
type UserLoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Paasword string `form:"paasword" json:"password" binding:"required"`
}

type userInfo struct {
	id       uint
	username string
	state    int16
}

func NewUserInfo(id uint, state int16, username string) *userInfo {
	return &userInfo{
		id:       id,
		state:    state,
		username: username,
	}
}

// user login
func (u *UserService) Login(userLogin UserLoginForm) *common.Response {
	user := &module.UserBasic{}
	user.getByUsername(userLogin.Paasword)
	if user.Username == "" {
		return common.ResError(http.StatusBadRequest, errors.New("用户名或密码错误"))
	}
	// todo 判断密码

	// todo 生成token

	// 返回登录信息
	return common.ResData(NewUserInfo(user.ID, user.State, user.Username))
}

// user logout
func (u *UserService) Logout(username string) {

}

// get userinfo
func (u *UserService) GetInfo(id int) *userInfo {
	return nil
}
