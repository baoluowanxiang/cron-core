package base

import (
	"errors"
	"fmt"
	"sync"
)

const (
	_ = iota*2 - 1
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

func (c *ClientOpt) NewOpt(key string, value interface{}) *Option {
	return &Option{key, value}
}

func (c *ClientOpt) SetOpt(opts ...*Option) *ClientOpt {
	for _, opt := range opts {
		c.Opts = append(c.Opts, opt)
	}
	return c
}

func (c *ClientOpt) GetOpt(key string) (interface{}, error) {
	for _, o := range c.Opts {
		if o.Key == key {
			return o.Value, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("key %s dosn't exist !", key))
}
