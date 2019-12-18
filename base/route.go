package base

import "github.com/gin-gonic/gin"

type Router interface {
	SetHttpRouter(*gin.Engine)
	SetTcpRouter(RouterMap)
}

type RouterMap interface {
	Put(name string, fn func())
	Get(name string) func()
}
