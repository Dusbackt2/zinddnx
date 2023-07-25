package main

import (
	"fmt"
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
		_, err := conn.Write([]byte("wo shi djj"))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}

		buf := make([]byte, 512)

		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}

		fmt.Printf(" server call back:%s,cnt=%d\n", buf, cnt)

		time.Sleep(1 * time.Second)

	}

}
