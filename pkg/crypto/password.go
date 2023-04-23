package password

import "golang.org/x/crypto/bcrypt"

type BcryptPassword struct{}

// 密码加密
func (b BcryptPassword) EncryptPassword(password string) (encrypt string, err error) {

	encryptPassword, erro := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if erro != nil {
		return "", err
	}
	return string(encryptPassword), nil
}

// 密码匹配校验
func (b BcryptPassword) MatchPassword(password string, encryptPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	return err == nil
}
