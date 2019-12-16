package entity

type Result struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Count int         `json:"count"`
}

const (
	CodeSuccess = 1
	CodeError   = -1
)
