package client_wrapper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// 注册信息
type registerInfo struct {
	ServiceName string `json:"service_name"`
	Token       string `json:"token"`
}

// 任务信息
type TaskInfo struct {
	Name   string      `json:"name"`
	ID     int         `json:"id"`
	Params interface{} `json:"params"`
}

type ConnectionWrapper struct {
	Resolver
}

//连接
func (cw *ConnectionWrapper) Connect() {

	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9600")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	log.Println(conn.LocalAddr().String() + " : Client connected!")

	// 初始化路由
	cw.Resolver.resolve()

	// 注册
	info := registerInfo{"tms", "aaaaaaaaaaaaaaaaaaaaaaa"}
	_info, _ := json.Marshal(info)
	_, _ = conn.Write([]byte(fmt.Sprintf("%s\n", string(_info))))
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			switch {
			case err == io.EOF: //读取完毕
				time.Sleep(100 * time.Millisecond)
				continue
			default: //连接断开了
				_ = conn.Close()
				goto end
			}
		}
		cw.resolve(str)
	}
end:
}

// 解析分发数据
func (cw *ConnectionWrapper) resolve(str string) {
	task := TaskInfo{}
	log.Print(str)
	err := json.Unmarshal([]byte(str), &task)
	if err != nil {
		log.Print("接收到的任务数据异常: ", err)
		goto end
	}
	cw.Resolver.execute(task)
end:
}
