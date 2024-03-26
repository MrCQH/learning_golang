package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 抽象了网络层和传输层
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	for {
		readString, _ := inputReader.ReadString('\n')
		readString = strings.Trim(readString, "\r\n")
		if readString == "quit" {
			return
		}
		fmt.Println("readString:", readString)
		conn.Write([]byte(readString))
	}
}
