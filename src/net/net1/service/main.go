package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error Listening", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		fmt.Printf("%s: 链接成功!!", conn.RemoteAddr())
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error Read", err.Error())
			return
		}
		fmt.Printf("Received data: %v\n", string(buf[:n]))
	}
}
