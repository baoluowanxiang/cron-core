package tcp

import (
	"crontab/base"
	"net"
	"sync"
	"time"
)

// 连接映射表
var connHashMap = ConnMap{}

var connChan = make(chan *connInfo)

// 连接服务hash表
type ConnMap map[string][]*connInfo

// 添加连接
func (c *ConnMap) addConn(conn *net.Conn, serviceName string, rp base.RouterMap) {
	var list []*connInfo
	var connMutex = sync.Mutex{}
	connMutex.Lock()
	list, ok := (*c)[serviceName]
	if !ok {
		list = []*connInfo{}
	}
	cInfo := &connInfo{}
	cInfo.Conn = conn
	cInfo.ServiceName = serviceName
	cInfo.RegisterTime = time.Now()
	cInfo.RouterMap = rp
	list = append(list, cInfo)
	(*c)[serviceName] = list
	connMutex.Unlock()
	connChan <- cInfo
}

// 监听服务连接
func (c *ConnMap) loop() {
	for cInfo := range connChan {
		go func() {
			// 地址会被覆盖
			info, err := cInfo.waitForMessage()
			if err != nil {
				c.deleteConn(info)
			}
		}()
	}
}

// 删除连接
func (c *ConnMap) deleteConn(info *connInfo) {
	list := (*c)[info.ServiceName]
	j := 0
	for _, info_c := range list {
		if info_c != info {
			list[j] = info_c
			j++
		}
	}
	(*c)[info.ServiceName] = list[:j]
}

func GetConnMap() ConnMap {
	return connHashMap
}
