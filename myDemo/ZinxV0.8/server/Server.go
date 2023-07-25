package main

import (
	"fmt"
	"myZinx8/ziface"
	"myZinx8/znet"
)

/**
基于zinx框架
*/

type PingRouter struct {
	znet.BaseRouter
}

func (y *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	fmt.Println("recv from client msgId:", request.GetMessageId(), ",data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(0, []byte("ping.ping.ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (y *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	fmt.Println("recv from client msgId:", request.GetMessageId(), ",data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("pong.pong.pong"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//创建一个server句柄,使用zinxApi
	server := znet.NewServer("dust5")
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})
	//启动Server
	server.Serve()
}
