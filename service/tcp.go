package service

import (
	"bufio"
	"crontab/base"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type TcpService struct {
	opt *base.ClientOpt
}

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
		//开启监听
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
	defer func() {
		_ = conn.Close()
	}()

	for {
		cmd, err := rw.ReadString('\n')
		cmd = strings.Trim(cmd, "\n ")
		log.Print(cmd)
		//runtime.Gosched()
		switch {
		case err == io.EOF:
			fmt.Println("读取完成.")
			return
		case err != nil:
			fmt.Println("读取出错")
			return
		}
	}
}

func (t *TcpService) loop() {

}
