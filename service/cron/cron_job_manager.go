package cron

import (
	"crontab/base"
	"crontab/entity"
	"crontab/job"
	repository "crontab/repository"
	runner2 "crontab/runner"
	"github.com/gin-gonic/gin"
)

var Manager cronJobManager

type cronJobManager struct {
	Cron base.CronService
	Tcp  base.TcpService
}

type AddJobRequest struct {
	Id     int    `json:"id" form:"id"`
	Schema string `json:"schema" form:"schema"`
	Route  string `json:"route" form:"route"`
}

func (t *cronJobManager) AddJob(ctx *gin.Context) {
	request := &AddJobRequest{}
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(502, "参数异常,"+err.Error())
		return
	}
	runner := &runner2.TcpRunner{Service: t.Tcp}
	j := &job.CronJob{}
	data := new(job.JobData)
	data.SetServiceName("tms")
	data.SetMessage(request.Route)
	j.Create(request.Id, request.Schema, data).SetRunner(runner)
	err = t.Cron.AddCron(j)
	if err != nil {
		ctx.JSON(502, "添加任务失败："+err.Error())
		return
	}
	ctx.JSON(200, "添加成功")
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
