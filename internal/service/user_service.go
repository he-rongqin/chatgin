package service

import (
	"errors"

	"github.com/sirupsen/logrus"
	"org.chatgin/internal/middleware"
	"org.chatgin/internal/module"
	"org.chatgin/internal/repository"

	password "org.chatgin/pkg/crypto"
)

type UserService struct {
}

// user register
func (u *UserService) Register(userRegister module.UserRegisterForm) error {
	// 校验手机号是否已经注册
	userRepositroy := &repository.UserRepository{}

	userBasci, err := userRepositroy.GetByUsername(userRegister.Phone)
	if err != nil {
		logrus.Errorf("查询用户发生错误：：%v", err)
		return errors.New("注册失败，系统错误")
	}
	if userBasci.Username != "" {
		return errors.New("注册失败，手机号已存在")
	}

	// 密码加密
	paaswordEncoder := &password.BcryptPassword{}
	password, err := paaswordEncoder.EncryptPassword(userRegister.Password)
	if err != nil {
		return errors.New("注册失败，系统异常")
	}
	userBasci.Passwrod = password
	userBasci.Username = userRegister.Phone
	userBasci.Account = userRegister.Phone
	userBasci.State = 1
	// 写入数据库
	return userRepositroy.Insert(userBasci)

}

// user login
func (u *UserService) Login(userLogin module.UserLoginForm) (user module.LoginUser, erro error) {
	var loginUser module.LoginUser
	if userLogin.Username == "" {
		return loginUser, errors.New("登录名不允许为空")
	}
	if userLogin.Paasword == "" {
		return loginUser, errors.New("登录密码不允许为空")
	}
	userRepositroy := &repository.UserRepository{}
	userBasic, err := userRepositroy.GetByUsername(userLogin.Username)
	if err != nil {
		logrus.Errorf("查询用户发生错误：：%v", err)
		return loginUser, errors.New("登录失败，系统错误。")
	}
	if user.Username == "" {
		return loginUser, errors.New("用户名或密码错误")
	}
	// 判断密码
	paaswordEncoder := &password.BcryptPassword{}
	if !paaswordEncoder.MatchPassword(userLogin.Paasword, userBasic.Passwrod) {
		return loginUser, errors.New("用户名或密码错误")

	}
	// todo 生成token
	tokenService := &middleware.TokenService{}
	claims := &middleware.JWTClaims{
		Custom: userBasic.ID,
	}
	token, erro := tokenService.Generate(*claims)
	if erro != nil {
		return loginUser, errors.New("登录失败，系统异常")
	}
	loginUser.Username = userBasic.Username
	loginUser.Id = userBasic.ID
	loginUser.State = userBasic.State
	loginUser.Token = token
	// 返回登录信息
	return loginUser, nil
}

// user logout
func (u *UserService) Logout(username string) bool {
	return true

}

// get userinfo
func (u *UserService) GetInfo(id uint) (userinfo *module.UserInfo, err error) {
	userRepository := &repository.UserRepository{}
	user, err := userRepository.GetById(id)
	if err != nil {
		return nil, errors.New("未找到相关用户数据")
	}
	return &module.UserInfo{
		Id:       user.ID,
		Username: user.Username,
		State:    user.State,
	}, nil
}
