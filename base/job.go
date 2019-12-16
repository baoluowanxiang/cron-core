package base

type Job interface {
	Run()
}

type JobData interface {
	SetServiceName(name string)

	GetServiceName() string

	SetMessage(msg string)

	GetMessage() string
}
