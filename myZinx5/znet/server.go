package znet

import (
	"fmt"
	"myZinx5/utils"
	"myZinx5/ziface"
	"net"
)

type Server struct {
	Name      string
	IPversion string
	IP        string
	port      int
	Router    ziface.IRouter
}

func (server *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP %s,Port %d,is starting\n", server.IP, server.port)
	fmt.Printf("[Zinx] Name %s, Version %s,MaxConn:%d,MaxPacketSize:%d\n", utils.GroubleObject.Name,
		utils.GroubleObject.Version, utils.GroubleObject.MaxConn, utils.GroubleObject.MaxPackageSize)
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
			dealConn := NewConnection(conn, cid, server.Router)
			cid++
			//启动当前的连接业务处理
			go dealConn.Start()
		}
	}()

}

func (server *Server) Stop() {
	//TODO将服务器资源或者状态，开辟的链接信息，停止或者回收

}

func (server *Server) Serve() {
	//异步
	server.Start()

	//TODO 启动服务器的时候的额外服务

	//阻塞状态
	select {}
}

func (server *Server) AddRouter(router ziface.IRouter) {
	server.Router = router
	fmt.Println("Add Router Success!!")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GroubleObject.Name,
		IPversion: "tcp4",
		IP:        utils.GroubleObject.Host,
		port:      utils.GroubleObject.TcpPort,
		Router:    nil,
	}
	return s
}
