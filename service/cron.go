package service

import (
	"crontab/base"
	"crontab/job"
	"github.com/robfig/cron/v3"
	"log"
)

type CronService struct {
	opt *base.ClientOpt
	cron *cron.Cron
	Ch chan *job.CronJob
	Sign chan string
}

func (crn *CronService) Start() {
	crn.Ch = make(chan *job.CronJob, 1000)
	crn.Sign = make(chan string, 2)
	defer func() {
		crn.opt.Wg.Done()
	}()
	crn.runListener()
	crn.runCron()
}


func (crn *CronService) SetOpt(opt *base.ClientOpt) {
	crn.opt = opt
}

func (crn *CronService) runCron() {
	opt := cron.WithParser(cron.NewParser(cron.Second|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow))
	crn.cron = cron.New(opt)
	crn.cron.Start()
}

func (crn *CronService) addCron(time string) {

}

func (crn *CronService) runListener() {
	go func() {
		for {
			select {
				case job := <-crn.Ch:
					crn.invokeJob(job)
				case <-crn.Sign:
					break;
			}
		}
	}()
}

func (crn *CronService) invokeJob(j *job.CronJob) {
	if j.OPC == job.OPC_ADD {
		entityId, err := crn.cron.AddJob(j.GetJobSchema(), j)
		if err != nil {
			log.Print(err)
			log.Print(entityId)
			return
		}
		j.SetState(job.JOB_STATE_ON)
		j.SetEntryId(entityId)
	} else if j.OPC == job.OPC_REMOVE {
		if j.GetState() != job.JOB_STATE_ON || j.GetEntryId() <=0 {
			return
		}
		crn.cron.Remove(j.GetEntryId())
	}
}