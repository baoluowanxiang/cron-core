package job

import (
	"crontab/base"
	runner2 "crontab/runner"
	"github.com/robfig/cron/v3"
	"log"
)

type CronJob struct {
	ID      int // 远程jobID
	Type    int // job 类型
	Schema  string
	Data    interface{}
	state   int          // 任务状态
	entryId cron.EntryID // entryId
	OPC     int          // opc 操作类型
}

const (
	_ = iota*2 - 1
	CRON_OPC_ADD
	CRON_OPC_REMOVE
)

const (
	_ = iota*2 - 1
	CRON_STATE_ON
	CRON_STATE_Off
)

const (
	_ = iota*2 - 1
	JOB_HTTP
	JOB_CMD
	JOB_TCP
)

func (j *CronJob) GetId() int {
	return j.ID
}

func (j *CronJob) SetId(id int) {
	j.ID = id
}

func (j *CronJob) GetJobType() int {
	return j.Type
}

func (j *CronJob) GetJobSchema() string {
	return j.Schema
}

func (j *CronJob) GetJobData() interface{} {
	return j.Data
}

func (j *CronJob) GetState() int {
	return j.state
}

func (j *CronJob) SetState(s int) {
	j.state = s
}

func (j *CronJob) GetEntryId() cron.EntryID {
	return j.entryId
}

func (j *CronJob) SetEntryId(id cron.EntryID) {
	j.entryId = id
}

func (j *CronJob) GetOpc() int {
	return j.OPC
}

func (j *CronJob) SetOpc(opc int) {
	j.OPC = opc
}

func (j *CronJob) Create(id int, jType int, schema string, data interface{}) *CronJob {
	j.ID = id
	j.Type = jType
	j.Data = data
	j.Schema = schema
	j.OPC = CRON_OPC_ADD
	return j
}

func (j *CronJob) Delete(id int) {

}

func (j *CronJob) Run() {
	log.Print(j.Data)
	var runner base.Runner
	if j.Type == JOB_TCP {
		runner = &runner2.TcpRunner{}
	} else if j.Type == JOB_CMD {
		runner = &runner2.HttpRunner{}
	} else {
		runner = &runner2.HttpRunner{}
	}
	go func() {
		runner.Run()
	}()
}
