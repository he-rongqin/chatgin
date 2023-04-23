package common

const (
	CONFIG_PREFIX            = "chatgin." // 配置前缀
	EVN_DEV                  = "dev"
	EVN_TEST                 = "test"
	EVN_PRE                  = "pre"
	EVN_PROD                 = "prod"
	DEFAULT_SERVER_PORT      = ":8080"
	DEFAULT_CONTEXT_PATH     = "/"
	DEFAULT_APP_NAME         = "app"
	HEADER_AUTHORIZATION_KEY = "Authorization"
	TOKEN_PREFIX             = "Bearer "
	NOT_AUTHENTICATION       = 401
	DEFAULT_PRIVATE_KEY      = "rongqin@chatgin#2023"
	DEFAULT_EXPIRES_TIMES    = 30
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    any    `json:"data"`
}

func ResSuccess() *Response {
	return &Response{
		Code:    200,
		Message: "success",
		Success: true,
		Data:    nil,
	}
}

func ResData(data any) *Response {
	return &Response{
		Code:    200,
		Message: "success",
		Success: true,
		Data:    data,
	}
}

func ResError(code int, e error) *Response {
	return &Response{
		Code:    code,
		Message: e.Error(),
		Success: false,
		Data:    nil,
	}
}

func ResErrorMsg(code int, msg string) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Success: false,
		Data:    nil,
	}
}
