package runner

import (
	"crontab/base"
)

type TcpRunner struct {
	base.BaseRunner
	Service base.TcpService
}

func (r *TcpRunner) Run(data base.JobData) {
	r.Service.Send(data)
}
