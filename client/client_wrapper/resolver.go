package client_wrapper

type Resolver interface {
	Resolve(name string, params interface{})
}
