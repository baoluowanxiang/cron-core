package http

import (
	"crontab/base"
	cron2 "crontab/service/cron"
	"crontab/service/cron/manager"
	tcp2 "crontab/service/tcp"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"log"
	"strings"
)

type Service struct {
	opt    *base.ClientOpt
	Client *cron2.CronService
	Router base.Router
}

func (h *Service) Start() error {
	go func() {
		err := h.exec()
		if err != nil {
			log.Fatal("启动http服务失败，失败原因：" + err.Error())
		}
		h.opt.Wg.Done()
	}()
	return nil
}

func (h *Service) SetOpt(opt *base.ClientOpt) {
	h.opt = opt
}

func (h *Service) Crock(client *cron2.CronService) *Service {
	h.Client = client
	return h
}

func (h *Service) WithRouter(rt base.Router) {
	h.Router = rt
}

func (h *Service) exec() error {
	portT, err := h.opt.GetOpt("port")
	if err != nil {
		return err
	}
	port, ok := portT.(string)
	if !ok {
		return errors.New("端口配置异常")
	}
	cronT, err := h.opt.GetOpt("cronService")
	cron, ok := cronT.(*cron2.CronService)
	if !ok {
		return errors.New("没有获取到cron")
	}

	tcpT, err := h.opt.GetOpt("tcpService")
	tcp, ok := tcpT.(*tcp2.Service)
	if !ok {
		return errors.New("没有获取到tcp服务")
	}

	h.Client = cron
	if h.Client == nil {
		return errors.New("请注入cron 服务")
	}

	r := gin.Default()
	manager.Init(cron, tcp)
	h.Router.SetHttpRouter(r)

	_ = r.Run(":" + strings.Trim(port, ":"))
	return nil
}
