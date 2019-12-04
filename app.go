package main

import (
	"crontab/base"
	"crontab/service"
	"sync"
	"time"
)

// 线程锁
func main() {
	client := &base.ClientOpt{
		Opts: []*base.Option{},
		Wg:   &sync.WaitGroup{},
	}
	cron := new(service.CronService)
	http := new(service.HttpService)
	tcp := new(service.TcpService)

	// 设置通道
	client.SetOpt(client.NewOpt("port", "3000"))
	client.SetOpt(client.NewOpt("cronService", cron))
	client.SetOpt(client.NewOpt("tcpService", tcp))

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
		var wg  = sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = srv.Start()
		}()
		wg.Wait()
	}
}
