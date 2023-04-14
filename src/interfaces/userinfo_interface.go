package interfaces

type UserInfo interface {
	// get username
	GetUsername() string

	// get userid
	GetId() uint

	// get user state
	GetState() int16

	// get usser token
	GetToken() string
}
