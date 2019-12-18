package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CrossXhr(c *gin.Context) {
	//请求方法
	method := c.Request.Method
	//请求头部
	origin := c.Request.Header.Get("Origin")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, X-Token, x-token, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token, Language, From, Scm-Token")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}
	// 处理请求
	c.Next() //  处理请求
}
