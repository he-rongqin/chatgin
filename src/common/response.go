package common

type Response struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	ErrorMsg error       `json:"errorMsg"`
	Data     interface{} `json:"data"`
}

func ResSuccess() *Response {
	return &Response{
		Code:     200,
		Message:  "success",
		ErrorMsg: nil,
		Data:     nil,
	}
}

func ResData(data interface{}) *Response {
	return &Response{
		Code:     200,
		Message:  "success",
		ErrorMsg: nil,
		Data:     data,
	}
}

func ResError(code int, e error) *Response {
	return &Response{
		Code:     code,
		ErrorMsg: e,
		Message:  "服务器错误，请稍后再试",
	}
}
