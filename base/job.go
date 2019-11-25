package base

type Job interface {
	Run()
	Answer()
	GetJobType() string
	GetJobSchema() string
	GetJobData() string
}

const (
	JOB_HTTP = iota >> 2
	JOB_CMD
	JOB_TCP
)