package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"org.chatgin/internal/config"
	"org.chatgin/internal/module"
)

type JWTClaims struct {
	Custom interface{}
	jwt.StandardClaims
}

type TokenService struct{}

func (t TokenService) Generate(claims JWTClaims) (userToken *module.UserToken, err error) {
	expireTime := time.Now().Add(time.Minute * time.Duration(config.TokenConfig.ExpiresTimes)).Unix() // 过期时间
	cl := JWTClaims{claims, jwt.StandardClaims{ExpiresAt: expireTime, Issuer: "rongqin@chatgin#2023"}}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, cl)
	str, err := token.SignedString([]byte(config.TokenConfig.PrivateKey))
	if err != nil {
		return nil, err
	}

	return &module.UserToken{
		AccessToekn:  str,
		ExpiresAt:    expireTime,
		RefreshToken: "",
	}, nil
}

// token 解析
func (t TokenService) Analyze(tokenStr string) (claims *JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.TokenConfig.PrivateKey), nil
	})

	if err != nil {
		logrus.Errorf("token 解析失败：%v error：%v", tokenStr, err)
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token 解析失败")

}
