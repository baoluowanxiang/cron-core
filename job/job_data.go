package job

import "crontab/base"

type JobData struct {
	name string
	data *base.JobParams
}

func (j *JobData) SetServiceName(name string) {
	j.name = name
}

func (j *JobData) GetServiceName() string {
	return j.name
}

func (j *JobData) SetData(data *base.JobParams) {
	j.data = data
}

func (j *JobData) GetData() *base.JobParams {
	return j.data
}
