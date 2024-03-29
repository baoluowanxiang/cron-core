package cron

import (
	"crontab/base"
	"crontab/job"
	"errors"
	"github.com/robfig/cron/v3"
	"log"
)

type Service struct {
	opt   *base.ClientOpt
	cron  *cron.Cron
	Ch    chan *job.CronJob
	Sign  chan int
	Table *job.JobHashTable
}

func (crn *Service) Start() error {
	crn.Ch = make(chan *job.CronJob, 1000)
	crn.Sign = make(chan int, 2)
	crn.Table = &job.JobHashTable{}
	defer func() {
		crn.opt.Wg.Done()
	}()
	crn.runListener()
	crn.runCron()
	return nil
}

func (crn *Service) SetOpt(opt *base.ClientOpt) {
	crn.opt = opt
}

func (crn *Service) WithRouter(rt base.Router) {

}

func (crn *Service) AddCron(cronJob base.Job) (err error) {
	switch cronJob.(type) {
	case *job.CronJob:
		var tmp interface{} = cronJob
		cJob := tmp.(*job.CronJob)
		cJob.OPC = job.CRON_OPC_ADD
		crn.Ch <- cJob
	default:
		err = errors.New("job 类型异常")
	}
	return err
}

func (crn *Service) StopCron(id int) error {
	cronJob := crn.Table.GetJob(id)
	if cronJob != nil {
		cronJob.OPC = job.CRON_OPC_REMOVE
		crn.Ch <- cronJob
		return nil
	}
	return errors.New("没有该任务，无法停止该任务")
}

func (crn *Service) runCron() {
	opt := cron.WithParser(cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow))
	crn.cron = cron.New(opt)
	crn.cron.Start()
}

func (crn *Service) runListener() {
	go func() {
		for {
			select {
			case j := <-crn.Ch:
				crn.invokeJob(j)
			case sign := <-crn.Sign:
				if sign == base.JOB_SIGN_STOP_CRON {
					return
				}
			}
		}
	}()
}

func (crn *Service) invokeJob(j *job.CronJob) {
	if j.OPC == job.CRON_OPC_ADD {
		if crn.Table.SetJob(j.ID, j) {
			entityId, err := crn.cron.AddJob(j.GetJobSchema(), j)
			if err != nil {
				log.Print(err)
				return
			}
			j.SetState(job.CRON_STATE_ON)
			j.SetEntryId(entityId)
		}
	} else if j.OPC == job.CRON_OPC_REMOVE {
		if j.GetState() != job.CRON_STATE_ON || j.GetEntryId() <= 0 {
			return
		}
		crn.cron.Remove(j.GetEntryId())
		j.SetState(job.CRON_STATE_Off)
		crn.Table.DelJob(j.ID)
	}
}
