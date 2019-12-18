package base

type Job interface {
	Run()
}

type JobParams struct {
	ID     int         `json:"id"`
	Name   string      `json:"name"`
	JType  int         `json:"j_type"`
	Params interface{} `json:"params"`
}

type JobData interface {
	SetServiceName(name string)

	GetServiceName() string

	SetData(data *JobParams)

	GetData() *JobParams
}
