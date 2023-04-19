package common

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
