package cron

import (
	"crontab/base"
	"crontab/entity"
	"crontab/job"
	repository "crontab/repository"
	runner2 "crontab/runner"
	"github.com/gin-gonic/gin"
	"time"
)

var Manager cronJobManager

type cronJobManager struct {
	Cron base.CronService
	Tcp  base.TcpService
}

type AddJobRequest struct {
	Id      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Schema  string `json:"schema" form:"schema"`
	Service string `json:"service_name" form:"service_name"`
	Job     string `json:"job" form:"job"`
}

func (t *cronJobManager) AddJob(ctx *gin.Context) {
	result := entity.Result{}
	request := &AddJobRequest{}
	err := ctx.ShouldBind(request)
	if err != nil {
		result.Msg = "添加任务失败：" + err.Error()
		result.Code = entity.CodeSuccess
		ctx.JSON(200, result)
		return
	}

	// 任务入库
	jobEntity := &entity.Job{}
	jobEntity.Name = request.Name
	jobEntity.Schema = request.Schema
	jobEntity.Job = request.Job
	jobEntity.ServiceName = request.Service
	jobEntity.CreateTime = time.Now()
	repos := repository.JobRepository{}
	repos.Save(jobEntity)

	// 任务入池
	runner := &runner2.TcpRunner{Service: t.Tcp}
	j := &job.CronJob{}
	data := new(job.JobData)
	data.SetServiceName(request.Service)
	data.SetData(&base.JobParams{jobEntity.ID, request.Job, 1, ""})
	j.Create(jobEntity.ID, request.Schema, data).SetRunner(runner)
	err = t.Cron.AddCron(j)

	if err != nil {
		result.Msg = "添加任务失败：" + err.Error()
		result.Code = entity.CodeSuccess
		ctx.JSON(502, result)
		return
	} else {
		result.Code = entity.CodeSuccess
		result.Msg = "添加成功"
		ctx.JSON(200, result)
		return
	}
}

func (t *cronJobManager) GetJobList(ctx *gin.Context) {
	result := entity.Result{}
	repos := repository.JobRepository{}
	list := repos.GetJobList()
	result.Code = entity.CodeSuccess
	result.Data = list
	result.Msg = "查询成功"
	ctx.JSON(200, result)
}

func Init(cron base.CronService, tcp base.TcpService) {
	Manager = cronJobManager{cron, tcp}
}
