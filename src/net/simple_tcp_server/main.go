package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const MaxRead = 25

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage : host port")
	}

	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("%s ", flag.Arg(i))
	}
	fmt.Println()

	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept")
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		readBytes := make([]byte, MaxRead+1)
		// 每次只读MaxRead长度, 多的长度分批发送
		n, err := conn.Read(readBytes[0:MaxRead])
		readBytes[MaxRead] = 0
		switch err {
		case nil:
			handleMsg(readBytes, n)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	err := conn.Close()
	fmt.Println("Closed connection: ", conn.RemoteAddr().String())
	checkError(err, "Close")
}

func handleMsg(readBytes []byte, n int) {
	if n > 0 {
		fmt.Print("<len: ", n, ". message: ")
		for i := 0; i < n-1; i++ {
			if readBytes[i] == 0 {
				break
			}
			fmt.Printf("%c", readBytes[i])
		}
		fmt.Print(".>\n")
	}
}

func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "ResolveTCPAddr")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP")
	return listener
}

func checkError(err error, info string) {
	if err != nil {
		fmt.Println("Error ", info, ": ", err.Error())
	}
}
