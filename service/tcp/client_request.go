package tcp

// client 注册请求
type ClientRegisterRequest struct {
	ServiceName string `json:"service_name"`
	Token       string `json:"token"`
}
