package job

import (
	"github.com/robfig/cron/v3"
	"log"
)

type CronJob struct {
	Type int // job 类型
	Schema string
	Data interface{}
	state int // 任务状态
	entryId cron.EntryID // entryid
	OPC	int // opc 操作类型
}

const (
	OPC_ADD = iota << 1
	OPC_REMOVE
)

const (
	JOB_STATE_Off = iota>>1
	JOB_STATE_ON
)

func(j *CronJob) Run() {
	log.Print(j.Data)
}

func(j *CronJob) Create(t int, schema string, data interface{}, o int) {
	j.Type = t
	j.Data = data
	j.Schema = schema
	j.OPC = o
}

func(j *CronJob) GetJobType() int {
	return j.Type
}

func(j *CronJob) GetJobSchema() string {
	return j.Schema
}

func(j *CronJob) GetJobData() interface{} {
	return j.Data
}

func(j *CronJob) GetState() int {
	return j.state
}

func (j *CronJob) SetState(s int) {
	j.state = s
}

func(j *CronJob) GetEntryId() cron.EntryID {
	return j.entryId
}

func (j *CronJob) SetEntryId(id cron.EntryID) {
	j.entryId = id
}

func(j *CronJob) GetOpc() int {
	return j.OPC
}