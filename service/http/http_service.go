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

type HttpService struct {
	opt    *base.ClientOpt
	Client *cron2.CronService
	Router base.Router
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

func (h *HttpService) SetOpt(opt *base.ClientOpt) {
	h.opt = opt
}

func (h *HttpService) Crock(client *cron2.CronService) *HttpService {
	h.Client = client
	return h
}

func (h *HttpService) WithRouter(rt base.Router) {
	h.Router = rt
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
	cron, ok := cron_t.(*cron2.CronService)
	if !ok {
		return errors.New("没有获取到cron")
	}

	tcp_t, err := h.opt.GetOpt("tcpService")
	tcp, ok := tcp_t.(*tcp2.TcpService)
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
