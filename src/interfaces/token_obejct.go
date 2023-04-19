package interfaces

type TokenObject[T any] interface {

	// token
	Token() string

	// 过期时间
	ExpiresAt() int64

	// 刷新token
	RefreshToken() string

	// 生成token
	Generate(claims T) error

	// 解析token
	Analyze(token string) (claims *T, err error)
}

type Claims interface {
}
