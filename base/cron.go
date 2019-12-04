package base

import (
	"crontab/job"
)

type CronService interface {
 	// 新增cron
	AddCron(cronJob *job.CronJob)
 	// 停止cron
 	StopCron(id int) error
}
