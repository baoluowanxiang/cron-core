package cron

import (
	"crontab/base"
	"crontab/job"
	runner2 "crontab/runner"
	"github.com/gin-gonic/gin"
)

var Manager CronJobManager

type CronJobManager struct {
	Cron base.CronService
	Tcp  base.TcpService
}

type AddJobRequest struct {
	Id     int    `json:"id" form:"id"`
	Schema string `json:"schema" form:"schema"`
	Data   string `json:"data" form:"data"`
}

func (t *CronJobManager) AddJob(ctx *gin.Context) {
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
	data.SetMessage(request.Data)
	j.Create(request.Id, request.Schema, data).SetRunner(runner)
	err = t.Cron.AddCron(j)
	if err != nil {
		ctx.JSON(502, "添加任务失败："+err.Error())
		return
	}
	ctx.JSON(200, "添加成功")
}

func Init(cron base.CronService, tcp base.TcpService) {
	Manager = CronJobManager{cron, tcp}
}
