package znet

import (
	"errors"
	"fmt"
	"io"
	"myZinx9/ziface"
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

	//消息通信
	msgChan chan []byte

	//该链接处理的方法router
	MsgHandle ziface.IMsgHandle

	//属于那个Server
	TcpServer ziface.IServer
}

// 初始化连接
func NewConnection(conn *net.TCPConn, connID uint32, handle ziface.IMsgHandle, tcpServer ziface.IServer) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		MsgHandle: handle,
		msgChan:   make(chan []byte),
		ExitChan:  make(chan bool, 1),
		TcpServer: tcpServer,
	}
	//将conn 加入manager
	c.TcpServer.GetConnManager().Add(c)
	return c
}

/*
*
用户发消息给客户端
*/
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Gorutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")
	for {

		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send Data error,", err)
				return
			}
		case <-c.ExitChan:
			//Reader 退出 Writer也退出
			return
		}

	}

}

func (c *Connection) StartReader() {
	fmt.Println("[Read Conn is running......]")
	defer fmt.Println("connID=", c.ConnID, "[Reader is exit],remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {

		//拆包捷豹对象
		pack := NewDataPack()
		bytes := make([]byte, pack.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), bytes)
		if err != nil {
			fmt.Println("unpark head error")
			break
		}

		msg, err := pack.Unpack(bytes)
		if err != nil {
			fmt.Println("server unpark err")
			return
		}
		if msg.GetMsgLen() > 0 {
			data := make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("read data err:", err)
				break
			}
			msg.SetData(data)
		}

		//得到当前conn的Request数据
		req := Request{
			conn: c,
			msg:  msg,
		}
		//结偶

		c.MsgHandle.SendMsgToTaskQueue(&req)
		//c.MsgHandle.DoMsgHandler(&req) 可以兜底

	}

}

// 提供一个SendMsg方法 将我们要发送给客户端的数据，先封包
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection is closed")
	}
	//封包
	pack := NewDataPack()
	msgPackage := NewMsgPackage(msgId, data)
	msg, err := pack.Pack(msgPackage)
	if err != nil {
		fmt.Println("error msgId:", msgId)
		return errors.New("pack error")
	}
	c.msgChan <- msg
	return nil
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()...ConnID = ", c.ConnID)
	//启动当前读数据的业务
	go c.StartReader()
	//TODO启动写数据业务
	go c.StartWriter()

	c.TcpServer.CallOnConnStart(c)

}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID=", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	c.TcpServer.CallOnConnStop(c)
	//调用关闭socket连接
	c.Conn.Close()

	//回收资源
	c.ExitChan <- true

	c.TcpServer.GetConnManager().Remove(c)
	close(c.msgChan)
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
