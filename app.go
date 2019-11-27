package main

import (
	"crontab/base"
	"crontab/service"
	"sync"
)

// 线程锁
func main() {
	client := base.ClientOpt{
		Opts: []*base.Option{},
		Wg:   &sync.WaitGroup{},
	}
	cron := new(service.CronService)
	client.SetOpt(client.NewOpt("port", "3000"))
	client.SetOpt(client.NewOpt("cron", cron))
	http := new(service.HttpService)
	registService(client, cron, http)
	client.Wg.Wait()
}

// 注册服务
func registService(client base.ClientOpt, services ...base.Service) {
	for _, srv := range services {
		client.Wg.Add(1)
		srv.SetOpt(&client)
		go func() {
			_ = srv.Start()
		}()
	}
}
