package tcp

import (
	"bufio"
	"crontab/base"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

type Service struct {
	opt       *base.ClientOpt
	Router    base.Router
	RouterMap base.RouterMap
}

func (ts *Service) Start() error {
	//var timeOut chan string
	go func() {
		err := ts.exec()
		if err != nil {
			log.Fatal("启动tcp服务失败，失败原因：", err.Error())
		}
	}()
	return nil
}

func (ts *Service) SetOpt(opt *base.ClientOpt) {
	ts.opt = opt
}

func (ts *Service) WithRouter(rt base.Router) {
	ts.Router = rt
}

func (ts *Service) Send(data base.JobData) {
	srvName := data.GetServiceName()
	conns, ok := connHashMap[srvName]
	if !ok {
		log.Print("未发现注册的服务")
		return
	}
	intN := len(conns)
	conn := conns[rand.Intn(intN)]
	dataBytes, _ := json.Marshal(data.GetData())
	_, _ = (*(*conn).Conn).Write([]byte(string(dataBytes) + "\n"))
}

func (ts *Service) exec() error {
	listener, err := net.Listen("tcp", "0.0.0.0:9600")
	if err != nil {
		log.Print("tcp service start error: ", err.Error())
		return err
	}
	ts.opt.Wg.Done()
	log.Print("tcp service start at port: 0.0.0.0:9600")

	// 注册路由
	router := ts.Router
	rmp := &RouterMap{}
	router.SetTcpRouter(rmp)
	ts.RouterMap = rmp

	// 等待连接认证
	go ts.waitForConnection(listener)

	// 等待消息
	go ts.loop()

	return nil
}

func (ts *Service) waitForConnection(listener net.Listener) {
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
			go ts.authenticate(cc)
		}
	}
}

// 接受数据， 验证身份
func (ts *Service) authenticate(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	done := make(chan int)
	ts.timeAccurate(conn, done)
	for {
		authString, err := rw.ReadString('\n')
		authString = strings.Trim(authString, "\n")
		if request, errc := ts.validateConnection(conn, authString); errc == nil {
			switch {
			case err == io.EOF: //客户端关闭了连接
				_ = conn.Close()
				goto end
			case err != nil: //连接断开了
				_ = conn.Close()
				goto end
			}
			ts.saveConnection(&conn, request)
			goto end
		} else {
			_, _ = conn.Write([]byte("unAuthenticate connection, " + errc.Error()))
			_ = conn.Close()
			goto end
		}
	}
end:
	{
		done <- 1
	}
}

// 5s之内没有完成认证，就算认证失败
func (ts *Service) timeAccurate(conn net.Conn, done chan int) {
	go func() {
		select {
		case <-time.After(5 * time.Second):
			log.Print("验证超时，链接断开")
			_ = conn.Close()
		case <-done:
			return
		}
	}()
}

// 验证身份
func (ts *Service) validateConnection(conn net.Conn, str string) (*ClientRegisterRequest, error) {
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

// 保存链接
func (ts *Service) saveConnection(conn *net.Conn, req *ClientRegisterRequest) {
	connHashMap.addConn(conn, req.ServiceName, ts.RouterMap)
}

func (ts *Service) loop() {
	connHashMap.loop()
}
