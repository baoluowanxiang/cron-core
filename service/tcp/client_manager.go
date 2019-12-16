package tcp

import (
	"crontab/entity"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type client struct {
	Name         string    `json:"name"`
	Host         string    `json:"host"`
	Schema       string    `json:"schema"`
	Status       int       `json:"status"`
	RegisterTime time.Time `json:"register_time"`
}

// 获取挂在在服务上的客户端列表
func GetServiceList(ctx *gin.Context) {
	result := entity.Result{}
	list := []client{}
	log.Print(connHashMap)
	for name, conns := range connHashMap {
		for _, info := range conns {
			clt := client{}
			clt.Name = name
			clt.Schema = (*(*info).Conn).RemoteAddr().Network()
			clt.Host = (*(*info).Conn).RemoteAddr().String()
			clt.Status = 1
			clt.RegisterTime = (*info).RegisterTime
			list = append(list, clt)
		}
	}
	result.Code = entity.CodeSuccess
	result.Data = list
	result.Count = len(list)
	ctx.JSON(200, result)
}
