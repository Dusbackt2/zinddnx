package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDatapack_Pack(t *testing.T) {
	//模拟服务起
	//

	listen, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("error 1")
	}

	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println("error 2")
			}
			go func(conn net.Conn) {
				//定义一个拆包的对象
				pack := NewDataPack()
				for {
					bytes := make([]byte, pack.GetHeadLen())
					_, err := io.ReadFull(conn, bytes)
					if err != nil {
						fmt.Println("read head error")
						break
					}
					unpack, err := pack.Unpack(bytes)
					if err != nil {
						fmt.Println("server unpark err")
						return
					}
					if unpack.GetMsgLen() > 0 {
						message := unpack.(*Message)
						message.Data = make([]byte, message.GetMsgLen())
						_, err := io.ReadFull(conn, message.Data)
						if err != nil {
							fmt.Println("server unpark data ", err)
							return
						}
						fmt.Println("--->Recv MsgID：", message.Id, ",dataLen=", message.DataLen, ",data=", string(message.Data))
					}
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:7777")

	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}

	dataPack := NewDataPack()

	m := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}

	pack, err := dataPack.Pack(m)
	if err != nil {
		fmt.Println("clent pack msg1 error", err)
		return
	}

	m2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'n', 'i', 'h', 'x', 'o', 'x', 'x'},
	}

	pack2, err := dataPack.Pack(m2)
	if err != nil {
		fmt.Println("clent pack msg1 error", err)
		return
	}
	pack = append(pack, pack2...)
	conn.Write(pack)

	select {}

}
