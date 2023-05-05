package common

const (
	CodeOK      = 200
	CodeFail    = 400
	MessageOK   = "ok"
	MessageFail = "fail"
)

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success() Resp {
	return Resp{
		Code:    CodeOK,
		Message: MessageOK,
		Data:    nil,
	}
}

func SuccessWith(message string, data any) Resp {
	if message == "" {
		message = MessageOK
	}
	return Resp{
		Code:    CodeOK,
		Message: message,
		Data:    data,
	}
}

func Fail() Resp {
	return Resp{
		Code:    CodeFail,
		Message: MessageFail,
		Data:    nil,
	}
}

func FailWith(message string, data any) Resp {
	if message == "" {
		message = MessageOK
	}
	return Resp{
		Code:    CodeFail,
		Message: message,
		Data:    data,
	}
}
