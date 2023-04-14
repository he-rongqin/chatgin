package service

import (
	"errors"

	"org.chatgin/src/authenticate"
	"org.chatgin/src/module"
)

type UserService struct {
}

// 用户登录表单结构
type UserLoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Paasword string `form:"paasword" json:"password" binding:"required"`
}

// 用户注册表单结构
type UserRegisterForm struct {
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Password string `form:"paasword"`
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

// user register
func (u *UserService) Register(userRegister UserRegisterForm) error {
	// 校验手机号是否已经注册
	userBasci := &module.UserBasic{}
	userBasci.GetByUsername(userRegister.Phone)
	if userBasci.Username != "" {
		return errors.New("注册失败，手机号已存在")
	}

	// 密码加密
	password, err := authenticate.EncryptPassword(userRegister.Password)
	if err != nil {
		return errors.New("注册失败，系统异常")
	}
	userBasci.Passwrod = password
	userBasci.Username = userRegister.Phone
	userBasci.State = 1
	// 写入数据库
	userBasci.Insert()
	return nil

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
	user.GetByUsername(userLogin.Paasword)
	if user.Username == "" {
		return nil, errors.New("用户名或密码错误")
	}
	// 判断密码
	if !authenticate.MatchPassword(userLogin.Paasword, user.Passwrod) {
		return nil, errors.New("用户名或密码错误")

	}
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
