package tcp

import (
	"net"
	"sync"
)

type Epoll struct {
	fd          int
	connections []net.Conn
	lock        sync.RWMutex
}
