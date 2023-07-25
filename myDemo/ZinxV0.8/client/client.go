package main

import (
	"fmt"
	"io"
	"myZinx8/znet"
	"net"
	"time"
)

func main() {
	fmt.Println("client start......")

	conn, err := net.Dial("tcp", "127.0.0.1:8369")
	if err != nil {
		fmt.Println("client start error,exit!")
	}

	for {
		//发送封包的msg
		pack := znet.NewDataPack()
		binary, err := pack.Pack(znet.NewMsgPackage(1, []byte("test message from djj")))
		if err != nil {
			fmt.Println("err;", err)
			return
		}
		if _, err := conn.Write(binary); err != nil {
			fmt.Println(" write err")
			return
		}

		headerbuf := make([]byte, pack.GetHeadLen())
		if _, err := io.ReadFull(conn, headerbuf); err != nil {
			fmt.Println("error  112121")
		}
		msghead, err := pack.Unpack(headerbuf)
		if err != nil {
			fmt.Println("msghead unpack error")
			break
		}
		if msghead.GetMsgLen() > 0 {
			message := msghead.(*znet.Message)
			message.Data = make([]byte, message.GetMsgLen())
			if _, err := io.ReadFull(conn, message.Data); err != nil {
				fmt.Println("unpack data error", err)
				break
			}
			fmt.Println("----->recv msgId:", message.Id, ",len:", message.DataLen, ",data:", string(message.Data))
		}

		time.Sleep(1 * time.Second)

	}

}
