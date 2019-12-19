package base

import "github.com/gin-gonic/gin"

type Router interface {
	SetHttpRouter(*gin.Engine)
	SetTcpRouter(RouterMap)
}

type RouteFn func()

type RouterMap interface {
	Put(string, RouteFn)
	Get(name string) RouteFn
}
