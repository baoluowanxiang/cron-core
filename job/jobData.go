package job

type JobData struct {
	name    string
	message string
}

func (j *JobData) SetServiceName(name string) {
	j.name = name
}

func (j *JobData) GetServiceName() string {
	return j.name
}

func (j *JobData) SetMessage(msg string) {
	j.message = msg
}

func (j *JobData) GetMessage() string {
	return j.message
}
