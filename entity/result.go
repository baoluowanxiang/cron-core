package entity

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	CodeSuccess = 1
	CodeError   = -1
)
