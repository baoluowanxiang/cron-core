package client_wrapper

// 路由表
var router = &Route{}

// handler
type Handler func(TaskInfo)

// 路由
type Route map[string]Handler

func (r *Route) Put(name string, handler Handler) {
	(*r)[name] = handler
}
func (r *Route) Get(name string) Handler {
	return (*r)[name]
}

// 解析器
type Resolver func(*Route)

func (r Resolver) resolve() {
	r(router)
}
func (r Resolver) execute(task TaskInfo) {
	if handle := router.Get(task.Name); handle != nil {
		go handle(task)
	}
}
