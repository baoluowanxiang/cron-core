package service

import (
	"crontab/base"
	job2 "crontab/job"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"strings"
)

type HttpService struct {
	opt    *base.ClientOpt
	Client *CronService
}

func (h *HttpService) Start() error {

	defer func() {
		h.opt.Wg.Done()
	}()

	port_t, err := h.opt.GetOpt("port")
	if err != nil {
		return err
	}
	port, ok := port_t.(string)
	if !ok {
		return errors.New("端口配置异常")
	}
	cron_t, err := h.opt.GetOpt("cron")
	cron, ok := cron_t.(*CronService)
	if !ok {
		return errors.New("没有获取到cron")
	}
	h.Client = cron
	if h.Client == nil {
		return errors.New("请注入cron 服务")
	}
	r := gin.Default()
	r.GET("/cron/add", h.AddJob)
	_ = r.Run(":" + strings.Trim(port, ":"))
	return nil
}

func (h *HttpService) SetOpt(opt *base.ClientOpt) {
	h.opt = opt
}

func (h *HttpService) Crock(client *CronService) *HttpService {
	h.Client = client
	return h
}

type AddJobRequest struct {
	Id     int         `json:"id"`
	Schema string      `json:"schema"`
	Data   interface{} `json:"data"`
}

func (h *HttpService) AddJob(ctx *gin.Context) {
	request := &AddJobRequest{}
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(502, "参数异常")
	}
	job := &job2.CronJob{}
	h.Client.AddCron(job.Create(request.Id, request.Schema, request.Data))
	ctx.JSON(200, "添加成功")
}
