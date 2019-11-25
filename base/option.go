package base

import "sync"

const (
	default_sign = iota << 2
	JOB_SIGN_STOP_CRON
)

type Option struct {
	Key   string
	Value interface{}
}

type ClientOpt struct {
	Opts []*Option
	Wg   *sync.WaitGroup
}

func (this *ClientOpt) SetOpt(opts ...*Option) *ClientOpt {
	for _, opt := range opts {
		this.Opts = append(this.Opts, opt)
	}
	return this
}
