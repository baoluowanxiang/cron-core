package base

type Service interface {
	SetOpt(opt *ClientOpt)
	Start() error
}

type CronService interface {
	// 新增cron
	AddCron(cronJob Job) error
	// 停止cron
	StopCron(id int) error
}

type TcpService interface {
	// 发送消息
	Send(data JobData)
}
