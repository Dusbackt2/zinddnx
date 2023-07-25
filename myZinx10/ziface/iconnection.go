package ziface

import "net"

type IConnection interface {
	//启动连接 让当前的连接准备开始工作
	Start()
	//停止连接 结束当前连接的工作
	Stop()
	//获取当前连接的绑定的socket conn
	GetTCPConnection() *net.TCPConn
	//获取当前连接模块的连接id
	GetConnID() uint32
	//获取对端的TCP状态 ip port
	RemoteAddr() net.Addr
	//发送数据，将数据发送给远程的客户端
	SendMsg(msgId uint32, data []byte) error

	SetProperty(key string, value interface{})

	GetProperty(key string) (interface{}, error)

	RemoveProrerty(key string)
}

// 定义一处理连接业务的方法   对端连接，当前处理数据的内容，内容长度
type HandleFunc func(*net.TCPConn, []byte, int) error
