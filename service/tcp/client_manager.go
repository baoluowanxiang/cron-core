package tcp

import (
	"crontab/entity"
	"github.com/gin-gonic/gin"
	"log"
)

type client struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Schema string `json:"schema"`
	Status int    `json:"status"`
}

// 获取挂在在服务上的客户端列表
func GetServiceList(ctx *gin.Context) {
	result := entity.Result{}
	list := []client{}
	log.Print(connHashMap)
	for name, conns := range connHashMap {
		for _, conn := range conns {
			clt := client{}
			clt.Name = name
			clt.Schema = (*conn).RemoteAddr().Network()
			clt.Host = (*conn).RemoteAddr().String()
			clt.Status = 1
			list = append(list, clt)
		}
	}
	result.Code = entity.CodeSuccess
	result.Data = list
	ctx.JSON(200, result)
}
