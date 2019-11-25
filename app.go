package main

import (
	"crontab/base"
	"crontab/job"
	"crontab/service"
	"sync"
	"time"
)

// 线程锁

func main() {
	client := base.ClientOpt{
		Opts: []*base.Option{},
		Wg:   &sync.WaitGroup{},
	}
	cron := new(service.CronService)
	registService(client, cron)
	client.Wg.Wait()
	s := new(job.CronJob)
	s.Create(1, "*/1 * * * * *", "hello, cron")
	var i int = 0
	for {
		time.Sleep(time.Second)
		if i < 3 {
			i++
			cron.AddCron(s)
		}
	}
}

// 注册服务
func registService(client base.ClientOpt, services ...base.Service) {
	for _, srv := range services {
		client.Wg.Add(1)
		srv.SetOpt(&client)
		go srv.Start()
	}
}
