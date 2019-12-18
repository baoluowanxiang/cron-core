package tcp

import (
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
func (c *ConnMap) addConn(conn *net.Conn, serviceName string) {
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
	list = append(list, cInfo)
	(*c)[serviceName] = list
	connMutex.Unlock()
	connChan <- cInfo
}

// 监听服务连接
func (c *ConnMap) loop() {
	for cInfo := range connChan {
		go cInfo.waitForMessage()
	}
}

func GetConnMap() ConnMap {
	return connHashMap
}
