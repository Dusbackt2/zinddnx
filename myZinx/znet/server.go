package znet

import (
	"fmt"
	"myZinx/ziface"
	"net"
)

type Server struct {
	Name      string
	IPversion string
	IP        string
	port      int
}

func (server *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP %s,Port %d,is starting\n", server.IP, server.port)

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
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			//客户端已经建立，做一些业务,做一个最基本的512字节长度的回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)

					if err != nil {
						fmt.Println("revc buf err", err)
						continue
					}

					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
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

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPversion: "tcp4",
		IP:        "0.0.0.0",
		port:      8369,
	}
	return s
}
