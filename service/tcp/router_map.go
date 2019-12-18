package tcp

type RouterMap map[string]func()

func (r *RouterMap) Put(name string, fn func()) {
	(*r)[name] = fn
}

func (r *RouterMap) Get(name string) func() {
	return (*r)[name]
}
