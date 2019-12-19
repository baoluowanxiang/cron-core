package main

import (
	"crontab/base"
	"crontab/route"
	"crontab/service/cron"
	"crontab/service/http"
	"crontab/service/tcp"
	"sync"
	"time"
)

// 线程锁
func main() {
	// 配置
	client := &base.ClientOpt{
		Opts: []*base.Option{},
		Wg:   &sync.WaitGroup{},
	}
	// 服务
	cronSrv := new(cron.Service)
	httpSrv := new(http.Service)
	tcpSrv := new(tcp.Service)
	// 设置通道
	client.SetOpt(client.NewOpt("port", "3000"))
	client.SetOpt(client.NewOpt("cronService", cronSrv))
	client.SetOpt(client.NewOpt("tcpService", tcpSrv))
	// 注册服务
	registService(client, cronSrv, httpSrv, tcpSrv)
	client.Wg.Wait()
	for {
		time.Sleep(time.Second)
	}
}

// 注册服务
func registService(opt *base.ClientOpt, services ...base.Service) {
	// 路由
	router := &route.Router{}

	for _, srv := range services {
		opt.Wg.Add(1)
		srv.SetOpt(opt)
		srv.WithRouter(router)
		var wg = sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = srv.Start()
		}()
		wg.Wait()
	}
}
