package interfaces

// 密码处理器接口
type PasswordEncoder interface {
	// 密码加密
	EncryptPassword(password string) (encrypt string, err error)

	// 密码匹配
	MatchPassword(password string, encryptPassword string) bool
}
