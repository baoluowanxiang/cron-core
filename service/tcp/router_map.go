package tcp

import "crontab/base"

type RouterMap map[string]func()

func (r *RouterMap) Put(name string, fn base.RouteFn) {
	(*r)[name] = fn
}

func (r *RouterMap) Get(name string) base.RouteFn {
	return (*r)[name]
}
