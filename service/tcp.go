package service

import (
	"crontab/base"
	"fmt"
	"log"
	"net"
)

type TcpService struct {
	opt    *base.ClientOpt
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


func (t *TcpService) exec() error {
	listener, err := net.Listen("tcp", "0.0.0.0:9600")
	if err != nil {
		log.Print("tcp service start error: ", err.Error())
		return err
	}
	t.opt.Wg.Done()
	log.Print("tcp service start at port: 0.0.0.0:9600")
	for {
		conn, err := listener.Accept() //开启监听
		if err != nil {
			fmt.Println("Accept is err!: ", err)
			continue
		}
		t.dealWithConnection(conn)
	}
}

func (t *TcpService) dealWithConnection(connection net.Conn) {
	log.Print("remote addr: ", connection.RemoteAddr())
}


func (t *TcpService) SetOpt(opt *base.ClientOpt) {
	t.opt = opt
}