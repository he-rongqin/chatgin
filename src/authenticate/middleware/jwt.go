package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"org.chatgin/src/config"
)

type JWTClaims struct {
	Custom interface{}
	jwt.StandardClaims
}

type Token struct {
	accessToken  string `json:"accessToken"`
	expiresAt    int64  `json:"expiresAt"`
	refreshToken string `json:"refreshToken"`
}

type TokenSericeJWT struct{}

func (t *Token) Generate(claims JWTClaims) error {
	expireTime := time.Now().Add(time.Minute * time.Duration(config.TokenConfig.ExpiresTimes)).Unix() // 过期时间
	cl := JWTClaims{claims, jwt.StandardClaims{ExpiresAt: expireTime, Issuer: "rongqin@chatgin#2023"}}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, cl)
	str, err := token.SignedString([]byte(config.TokenConfig.PrivateKey))
	if err != nil {
		return err
	}
	t.accessToken = str
	t.expiresAt = expireTime
	// t.RefreshToken
	return nil
}

func (t *Token) Token() string {
	return t.accessToken
}

func (t *Token) ExpiresAt() int64 {
	return t.expiresAt
}

func (t *Token) RefreshToken() string {
	return ""
}

// token 解析
func (t *Token) Analyze(tokenStr string) (claims *JWTClaims, err error) {
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
