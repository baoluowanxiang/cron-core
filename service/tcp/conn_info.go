package tcp

import (
	"bufio"
	"crontab/base"
	"errors"
	"io"
	"net"
	"strings"
	"time"
)

// 连接信息
type connInfo struct {
	Conn         *net.Conn
	ServiceName  string
	RegisterTime time.Time
	RouterMap    base.RouterMap
}

func (c *connInfo) waitForMessage() error {
	rw := bufio.NewReadWriter(bufio.NewReader(*c.Conn), bufio.NewWriter(*c.Conn))
	for {
		askStr, err := rw.ReadString('\n')
		askStr = strings.Trim(askStr, "\n")
		switch {
		case err == io.EOF: //客户端关闭了连接
			return errors.New("连接已关闭")
		case err != nil: //连接断开了
			return errors.New("连接已断开")
		}
	}
}
