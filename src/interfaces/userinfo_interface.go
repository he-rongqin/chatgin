package interfaces

type LoginUser[T any] interface {
	// get username
	GetUsername() string

	// get userid
	GetId() uint

	// get user state
	GetState() int16

	// get usser token
	GetToken() TokenObject[T]
}
