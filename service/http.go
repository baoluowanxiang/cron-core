package service

import (
	"crontab/base"
	"crontab/route"
	cron2 "crontab/service/cron"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"log"
	"strings"
)

type HttpService struct {
	opt    *base.ClientOpt
	Client *CronService
}

func (h *HttpService) Start() error {
	go func() {
		err := h.exec()
		if err != nil {
			log.Fatal("启动http服务失败，失败原因：" + err.Error())
		}
		h.opt.Wg.Done()
	}()
	return nil
}

func (h *HttpService) exec() error {
	port_t, err := h.opt.GetOpt("port")
	if err != nil {
		return err
	}
	port, ok := port_t.(string)
	if !ok {
		return errors.New("端口配置异常")
	}
	cron_t, err := h.opt.GetOpt("cronService")
	cron, ok := cron_t.(*CronService)
	if !ok {
		return errors.New("没有获取到cron")
	}

	tcp_t, err := h.opt.GetOpt("tcpService")
	tcp, ok := tcp_t.(*TcpService)
	if !ok {
		return errors.New("没有获取到tcp服务")
	}

	h.Client = cron
	if h.Client == nil {
		return errors.New("请注入cron 服务")
	}
	r := gin.Default()
	cron2.Init(cron, tcp)
	route.SetRouter(r)
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
