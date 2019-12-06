package base

type CronService interface {
	// 新增cron
	AddCron(cronJob Job) error
	// 停止cron
	StopCron(id int) error
}
