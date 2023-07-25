package znet

import "myZinx3/ziface"

type Request struct {
	//已经和客户端建立好的连接 Conn
	conn ziface.IConnection
	//客户端请求的数据
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
