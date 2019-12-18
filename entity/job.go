package entity

import "time"

type Job struct {
	ID              int       `gorm:"column:id" json:"id"`
	Name            string    `gorm:"column:name" json:"name"`
	Schema          string    `gorm:"column:crontab" json:"schema"`
	ServiceName     string    `gorm:"column:service_name" json:"service_name"`
	Job             string    `gorm:"column:job" json:"job"`
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	LastExecuteTime time.Time `gorm:"column:last_execute_time" json:"last_execute_time"`
}

func (j *Job) TableName() string {
	return "cron_job"
}
