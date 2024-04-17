package global

// Msg 响应消息
// type Msg struct {
// 	Code int         `json:"code"`
// 	Msg  string      `json:"msg"`
// 	Data interface{} `json:"data"`
// }

// 业务码
const (
	SuccessCode = 0
	FailCode    = 1
)

// 业务消息
const (
	SuccessMsg  = "success"
	FailMsg     = "fail"
	TokenExpire = "token expire"
)
