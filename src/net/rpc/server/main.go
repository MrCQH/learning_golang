package main

import (
	"learning_golang/net/rpc"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	calc := new(rpc_objects.Args)
	err := rpc.Register(calc)
	if err != nil {
		log.Fatal("RPC Register failure", err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", "localhost:1234")
	defer listener.Close()
	if err != nil {
		log.Fatal("Start RPC-server-listen error:", err)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}
