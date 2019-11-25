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
	cron := new(service.CronService);
	registService(client, cron)
	client.Wg.Wait()
	var i int = 0
	j := new(job.CronJob)
	j.Create(1, "*/1 * * * * *", "hello,world", job.OPC_ADD )
	for{
		time.Sleep(time.Second)
		i++
		if i==10 {
			cron.Ch<-j
		}
		if i == 20 {
			j.OPC = job.OPC_REMOVE
			cron.Ch<-j
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

