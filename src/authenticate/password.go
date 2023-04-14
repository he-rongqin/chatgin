package authenticate

import "golang.org/x/crypto/bcrypt"

// 密码加密
func EncryptPassword(password string) (encrypt string, err error) {

	encryptPassword, erro := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if erro != nil {
		return "", err
	}
	return string(encryptPassword), nil
}

// 密码匹配校验
func MatchPassword(password string, encryptPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	return err == nil
}
