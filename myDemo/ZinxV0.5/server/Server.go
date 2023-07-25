package main

import (
	"fmt"
	"myZinx5/ziface"
	"myZinx5/znet"
)

/**
基于zinx框架
*/

type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	fmt.Println("recv from client msgId:", request.GetMessageId(), ",data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping.ping.ping"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//创建一个server句柄,使用zinxApi
	server := znet.NewServer("dust5")
	server.AddRouter(&PingRouter{})
	//启动Server
	server.Serve()
}
