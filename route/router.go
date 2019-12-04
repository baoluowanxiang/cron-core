package route

import (
	"crontab/service/cron"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {

	// 新增任务
	r.POST("/cron/add", cron.Manager.AddJob);

	// 查询执行端列表


}