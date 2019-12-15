package tcp

import (
	"crontab/entity"
	"github.com/gin-gonic/gin"
)

type client struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Status int    `json:"status"`
}

// 获取挂在在服务上的客户端列表
func GetServiceList(ctx *gin.Context) {
	result := entity.Result{}
	list := []client{}
	for name, conns := range connHashMap {
		for _, conn := range conns {
			clt := client{}
			clt.Name = name
			clt.Host = (*conn).LocalAddr().Network()
			clt.Status = 1
		}
	}
	result.Code = entity.CodeSuccess
	result.Data = list
	ctx.JSON(200, result)
}
