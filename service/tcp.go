package service

import (
	"bufio"
	"crontab/base"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

type TcpService struct {
	opt       *base.ClientOpt
	connMutex sync.Mutex
}

// 连接服务hash表
type connMap map[string][]*net.Conn

// client 注册请求
type ClientRegisterRequest struct {
	ServiceName string `json:"service_name"`
	Token       string `json:"token"`
}

// 连接映射表
var connHashMap = connMap{}

// 连接通道
var connChannel = make(chan *net.Conn)

func (t *TcpService) Start() error {
	go func() {
		err := t.exec()
		if err != nil {
			log.Fatal("启动tcp服务失败，失败原因：", err.Error())
		}
	}()
	return nil
}

func (t *TcpService) SetOpt(opt *base.ClientOpt) {
	t.opt = opt
	t.connMutex = sync.Mutex{}
}

func (t *TcpService) Send(data base.JobData) {
	log.Print("tcp:", data)
}

func (t *TcpService) exec() error {
	listener, err := net.Listen("tcp", "0.0.0.0:9600")
	if err != nil {
		log.Print("tcp service start error: ", err.Error())
		return err
	}
	t.opt.Wg.Done()
	log.Print("tcp service start at port: 0.0.0.0:9600")
	// 等待连接认证
	go t.waitForConnection(listener)
	// 等待消息
	go t.loop()

	return nil
}

func (t *TcpService) waitForConnection(listener net.Listener) {
	var connChan = make(chan net.Conn)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept is err!: ", err)
				continue
			}
			connChan <- conn
		}
	}()
	for {
		select {
		case cc := <-connChan:
			go t.authenticate(cc)
		}
	}
}

func (t *TcpService) authenticate(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		authString, err := rw.ReadString('\n')
		authString = strings.Trim(authString, "\n ")
		log.Print(authString, err)
		if data, errc := t.validateConnection(conn, authString); errc == nil {
			switch {
			case err == io.EOF: //客户端关闭了连接
				_ = conn.Close()
				return
			case err != nil: //连接断开了
				_ = conn.Close()
				return
			}
			t.saveConnection(conn, data)
			return
		} else {
			_, _ = conn.Write([]byte("unAuthenticate connection, " + errc.Error()))
			_ = conn.Close()
			return
		}
	}
}

func (t *TcpService) validateConnection(str string) (*ClientRegisterRequest, error) {
	request := &ClientRegisterRequest{}
	err := json.Unmarshal([]byte(str), request)
	if err != nil {
		return nil, errors.New("数据结构异常")
	}
	if request.ServiceName == "" || request.Token == "" {
		return nil, errors.New("参数异常")
	}
	return request, nil
}

func (t *TcpService) saveConnection(conn net.Conn, req *ClientRegisterRequest) {
	var list []*net.Conn
	t.connMutex.Lock()
	list, ok := connHashMap[req.ServiceName]
	if !ok {
		list = []*net.Conn{}
	}
	list = append(list, &conn)
	connHashMap[req.ServiceName] = list
	t.connMutex.Unlock()
}

func (t *TcpService) loop() {

}
