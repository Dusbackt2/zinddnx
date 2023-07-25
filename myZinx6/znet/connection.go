package znet

import (
	"errors"
	"fmt"
	"io"
	"myZinx6/ziface"
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
	MsgHandle ziface.IMsgHandle
}

// 初始化连接
func NewConnection(conn *net.TCPConn, connID uint32, handle ziface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		MsgHandle: handle,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Read Conn is running......")
	defer fmt.Println("connID=", c.ConnID, "Reader is exit,remote addr is", c.RemoteAddr().String())
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

		c.MsgHandle.DoMsgHandler(&req)
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
	_, err = c.GetTCPConnection().Write(msg)
	if err != nil {
		fmt.Println(" write error :", err)
		return errors.New("write buf error")
	}
	return nil
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
