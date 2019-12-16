package main

import (
	"crontab/base"
	cron2 "crontab/service/cron"
	http2 "crontab/service/http"
	tcp2 "crontab/service/tcp"
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
	cron := new(cron2.CronService)
	http := new(http2.HttpService)
	tcp := new(tcp2.TcpService)

	// 设置通道
	client.SetOpt(client.NewOpt("port", "3000"))
	client.SetOpt(client.NewOpt("cronService", cron))
	client.SetOpt(client.NewOpt("tcpService", tcp))

	// 注册服务
	registService(client, cron, http, tcp)
	client.Wg.Wait()
	for {
		time.Sleep(time.Second)
	}
}

// 注册服务
func registService(client *base.ClientOpt, services ...base.Service) {
	for _, srv := range services {
		client.Wg.Add(1)
		srv.SetOpt(client)
		var wg = sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = srv.Start()
		}()
		wg.Wait()
	}
}
