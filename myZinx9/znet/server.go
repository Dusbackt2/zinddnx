package znet

import (
	"fmt"
	"myZinx9/utils"
	"myZinx9/ziface"
	"net"
)

type Server struct {
	Name        string
	IPversion   string
	IP          string
	port        int
	MsgHandle   ziface.IMsgHandle
	ConnManager ziface.IConnManager
	OnConnStart func(conn ziface.IConnection)
	OnConnEnd   func(conn ziface.IConnection)
}

func (server *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP %s,Port %d,is starting\n", server.IP, server.port)
	fmt.Printf("[Zinx] Name %s, Version %s,MaxConn:%d,MaxPacketSize:%d\n", utils.GroubleObject.Name,
		utils.GroubleObject.Version, utils.GroubleObject.MaxConn, utils.GroubleObject.MaxPackageSize)
	server.MsgHandle.StartWorkerPool()
	go func() {
		//1.获取TCP的Addr
		addr, err := net.ResolveTCPAddr(server.IPversion, fmt.Sprintf("%s:%d", server.IP, server.port))
		if err != nil {
			fmt.Println("resolve tcp addt error :", err)
			return
		}
		//2.监听服务器地址
		listener, err := net.ListenTCP(server.IPversion, addr)
		if err != nil {
			fmt.Println("listen", server.IPversion, "err:", err)
			return
		}
		//3.阻塞等待客户端链接，处理客户端链接业务（读写）
		fmt.Println("start Zinx server succ,", server.Name, " succ,Listenning...")
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//设置最大连接个数的判
			if server.ConnManager.Len() >= 2 {
				//TODO 给客户端响应错误包
				fmt.Println("[max conn limit]=================================================================================")
				conn.Close()
				continue
			}

			dealConn := NewConnection(conn, cid, server.MsgHandle, server)
			cid++
			//启动当前的连接业务处理
			go dealConn.Start()
		}
	}()

}

func (server *Server) Stop() {
	//TODO将服务器资源或者状态，开辟的链接信息，停止或者回收
	server.ConnManager.ClearConn()
}

func (server *Server) Serve() {
	//异步
	server.Start()

	//TODO 启动服务器的时候的额外服务

	//阻塞状态
	select {}
}

func (server *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	server.MsgHandle.AddRouter(msgId, router)
	fmt.Println("Add Router Success!!")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:        utils.GroubleObject.Name,
		IPversion:   "tcp4",
		IP:          utils.GroubleObject.Host,
		port:        utils.GroubleObject.TcpPort,
		MsgHandle:   NewMsgHandle(),
		ConnManager: NewConnManager(),
	}
	return s
}

func (server *Server) GetConnManager() ziface.IConnManager {
	return server.ConnManager
}

func (server *Server) SetOnConnStart(hookfunc func(connection ziface.IConnection)) {
	server.OnConnStart = hookfunc
}

func (server *Server) SetOnConnStop(hookfunc func(connection ziface.IConnection)) {
	server.OnConnEnd = hookfunc
}

func (server *Server) CallOnConnStart(connection ziface.IConnection) {
	if server.OnConnStart != nil {
		server.OnConnStart(connection)
	}
}

func (server *Server) CallOnConnStop(connection ziface.IConnection) {
	if server.OnConnEnd != nil {
		server.OnConnEnd(connection)
	}
}
