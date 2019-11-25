package base

import "sync"

type Option struct {
	Key string
	Value interface{}
}

type ClientOpt struct {
	Opts []*Option
	Wg *sync.WaitGroup
}

func (this *ClientOpt) SetOpt(opts ...*Option) *ClientOpt {
	for _, opt := range opts {
		this.Opts = append(this.Opts, opt)
	}
	return this
}
