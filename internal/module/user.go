package module

type LoginUser struct {
	Id       uint       `json:"id"`
	Username string     `json:"username"`
	State    int16      `json:"state"`
	Token    *UserToken `json:"token"`
}

type UserToken struct {
	AccessToekn  string `json:"accessToken"`
	ExpiresAt    int64  `json:"expiresAt"`
	RefreshToken string `jsron:"refreshToken"`
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
	Id       uint   `json:"id"`
	Username string `json:"username"`
	State    int16  `json:"state"`
}
