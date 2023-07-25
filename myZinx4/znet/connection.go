package znet

import (
	"fmt"
	"myZinx4/utils"
	"myZinx4/ziface"
	"net"
)

/**
当前连接模块
*/

type Connection struct {
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//连接的ID
	ConnID uint32
	//当前连接的状态
	isClosed bool
	//当前连接所绑定的处理业务方法API
	handleAPI ziface.HandleFunc
	//告知当前连接已经退出的/停止channel
	ExitChan chan bool
	//该链接处理的方法router
	Router ziface.IRouter
}

// 初始化连接
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Read Conn is running......")
	defer fmt.Println("connID=", c.ConnID, "Reader is exit,remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中 512字节
		buf := make([]byte, utils.GroubleObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}
		//得到当前conn的Request数据
		req := Request{
			conn: c,
			data: buf,
		}

		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}

}

func (c *Connection) Start() {
	fmt.Println("Conn Start()...ConnID = ", c.ConnID)
	//启动当前读数据的业务
	go c.StartReader()
	//TODO启动写数据业务

}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID=", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	//调用关闭socket连接
	c.Conn.Close()
	//回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
