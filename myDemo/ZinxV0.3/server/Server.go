package main

import (
	"fmt"
	"myZinx3/ziface"
	"myZinx3/znet"
)

/**
基于zinx框架
*/

type PingRouter struct {
	znet.BaseRouter
}

// TestPreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	connection := request.GetConnection().GetTCPConnection()
	_, err := connection.Write([]byte("ddduuusssttt333"))
	if err != nil {
		fmt.Println("cal back ping error")
	}
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	connection := request.GetConnection().GetTCPConnection()
	_, err := connection.Write([]byte("ping...ping...ping..."))
	if err != nil {
		fmt.Println("cal back ping error")
	}
}

func main() {
	//创建一个server句柄,使用zinxApi
	server := znet.NewServer("dust3")
	server.AddRouter(&PingRouter{})
	//启动Server
	server.Serve()
}
