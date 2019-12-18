package repository

import (
	"crontab/entity"
	"crontab/orm"
)

type JobRepository struct {
}

// 保存任务
func (j *JobRepository) Save(job *entity.Job) {
	orm.GlobalDb.Table((&entity.Job{}).TableName()).Save(job)
}

// 获取任务列表
func (j *JobRepository) GetJobList() []*entity.Job {
	var jobList []*entity.Job
	query := orm.GlobalDb.Table((&entity.Job{}).TableName())
	query = query.Find(&jobList)
	return jobList
}
