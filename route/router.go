package route

import (
	"crontab/base"
	"crontab/route/middleware"
	mgr "crontab/service/cron/manager"
	"crontab/service/tcp/manager"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

// http 路由
func (rt *Router) SetHttpRouter(r *gin.Engine) {
	// 跨域
	r.Use(middleware.CrossXhr)
	// cron 管理
	crt := r.Group("/cron/")
	{
		// 新增任务
		crt.POST("job/add", mgr.Manager.AddJob)
		// 获取任务列表
		crt.GET("job/list", mgr.Manager.GetJobList)
		// 查询执行端列表
		crt.GET("agent/list", manager.GetServiceList)
	}
}

// tcp 路由
func (rt *Router) SetTcpRouter(r base.RouterMap) {
	// 回传执行日志
	r.Put("log", func() {})
	// 心跳检测
	r.Put("ping", func() {})
}
