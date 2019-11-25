package job

import (
	"github.com/robfig/cron/v3"
	"log"
)

type CronJob struct {
	ID      int // 远程jobID
	Type    int // job 类型
	Schema  string
	Data    interface{}
	state   int          // 任务状态
	entryId cron.EntryID // entryid
	OPC     int          // opc 操作类型
}

const (
	CRON_OPC_ADD = iota << 1
	CRON_OPC_REMOVE
)

const (
	CRON_STATE_Off = iota << 1
	CRON_STATE_ON
)

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

func (j *CronJob) Create(id int, schema string, data interface{}) {
	j.ID = id
	j.Data = data
	j.Schema = schema
	j.OPC = CRON_OPC_ADD
}

func (j *CronJob) Delete(id int) {

}

func (j *CronJob) Run() {

	log.Print(j.Data)

}
