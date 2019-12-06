package cron

import (
	"crontab/base"
	"crontab/job"
	"github.com/gin-gonic/gin"
	"log"
)

var Manager CronJobManager

type CronJobManager struct {
	Client base.CronService
}

type AddJobRequest struct {
	Id     int    `json:"id" form:"id"`
	Schema string `json:"schema" form:"schema"`
	Data   string `json:"data" form:"data"`
}

func (h *CronJobManager) AddJob(ctx *gin.Context) {
	request := &AddJobRequest{}
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(502, "参数异常,"+err.Error())
		return
	}
	j := &job.CronJob{}
	log.Print("request:", request)
	h.Client.AddCron(j.Create(request.Id, job.JOB_TCP, request.Schema, request.Data))
	ctx.JSON(200, "添加成功")
}

func Init(cron base.CronService) {
	Manager = CronJobManager{cron}
}
