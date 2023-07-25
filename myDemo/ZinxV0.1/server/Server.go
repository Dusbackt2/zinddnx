package main

import "myZinx/znet"

/**
基于zinx框架
*/

func main() {
	//创建一个server句柄,使用zinxApi
	server := znet.NewServer("dust")
	//启动Server
	server.Serve()
}
