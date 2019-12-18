package tcp

import (
	"net"
	"time"
)

// 连接信息
type connInfo struct {
	Conn         *net.Conn
	ServiceName  string
	RegisterTime time.Time
}

func (c *connInfo) waitForMessage() {

}
