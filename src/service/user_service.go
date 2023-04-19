package service

import (
	"errors"

	"github.com/sirupsen/logrus"
	"org.chatgin/src/authenticate"
	"org.chatgin/src/authenticate/middleware"
	"org.chatgin/src/interfaces"
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

type LoginUser struct {
	id       uint                                         `json:"id"`
	username string                                       `json:"username"`
	state    int16                                        `json:"state"`
	token    interfaces.TokenObject[middleware.JWTClaims] `json:"token"`
}

type UserInfo struct {
	Id       uint
	Username string
	State    int16
}

func (userinfo *LoginUser) GetUsername() string {
	return userinfo.username
}

func (userinfo *LoginUser) GetId() uint {
	return userinfo.id
}

func (userinfo *LoginUser) GetState() int16 {
	return userinfo.state
}

func (userinfo *LoginUser) GetToken() interfaces.TokenObject[middleware.JWTClaims] {
	return userinfo.token
}

func NewLoginUserInfo(id uint, state int16, username string, token interfaces.TokenObject[middleware.JWTClaims]) *LoginUser {
	return &LoginUser{
		id:       id,
		state:    state,
		username: username,
		token:    token,
	}
}

// user register
func (u *UserService) Register(userRegister UserRegisterForm) error {
	// 校验手机号是否已经注册
	userBasci := &module.UserBasic{}
	if err := userBasci.GetByUsername(userRegister.Phone); err != nil {
		logrus.Errorf("查询用户发生错误：：%v", err)
		return errors.New("注册失败，系统错误")
	}
	if userBasci.Username != "" {
		return errors.New("注册失败，手机号已存在")
	}

	// 密码加密
	paaswordEncoder := &authenticate.BcryptPassword{}
	password, err := paaswordEncoder.EncryptPassword(userRegister.Password)
	if err != nil {
		return errors.New("注册失败，系统异常")
	}
	userBasci.Passwrod = password
	userBasci.Username = userRegister.Phone
	userBasci.Account = userRegister.Phone
	userBasci.State = 1
	// 写入数据库
	return userBasci.Insert()

}

// user login
func (u *UserService) Login(userLogin UserLoginForm) (userInfo interfaces.LoginUser[middleware.JWTClaims], erro error) {
	if userLogin.Username == "" {
		return nil, errors.New("登录名不允许为空")
	}
	if userLogin.Paasword == "" {
		return nil, errors.New("登录密码不允许为空")
	}
	user := &module.UserBasic{}
	if err := user.GetByUsername(userLogin.Username); err != nil {
		logrus.Errorf("查询用户发生错误：：%v", err)
		return nil, errors.New("登录失败，系统错误。")
	}
	if user.Username == "" {
		return nil, errors.New("用户名或密码错误")
	}
	// 判断密码
	paaswordEncoder := &authenticate.BcryptPassword{}
	if !paaswordEncoder.MatchPassword(userLogin.Paasword, user.Passwrod) {
		return nil, errors.New("用户名或密码错误")

	}
	// todo 生成token
	token := &middleware.Token{}
	claims := &middleware.JWTClaims{
		Custom: user.ID,
	}
	erro = token.Generate(*claims)
	if erro != nil {
		return nil, errors.New("登录失败，系统异常")
	}
	// 返回登录信息
	return NewLoginUserInfo(user.ID, user.State, user.Username, token), nil
}

// user logout
func (u *UserService) Logout(username string) bool {
	return true

}

// get userinfo
func (u *UserService) GetInfo(id uint) (userinfo *UserInfo, err error) {
	user := &module.UserBasic{}
	if err := user.GetById(id); err != nil {
		return nil, errors.New("未找到相关用户数据")
	}
	return &UserInfo{
		Id:       user.ID,
		Username: user.Username,
		State:    user.State,
	}, nil
}
