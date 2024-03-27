package main

import (
	"fmt"
	"learning_golang/net/rpc"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	args := &rpc_objects.Args{N: 7, M: 8}

	// 同步
	var res int
	sync("Args.Multiply", args, &res, client)
	fmt.Println("sync result: ", res)
	var asyncRes int

	// 异步等待
	//call := async1("Args.Multiply", args, &asyncRes, client)
	//fmt.Println("async result: ", asyncRes, " ", call)

	// 异步不等待 + sleep模拟结果返回
	async2("Args.Multiply", args, &asyncRes, client)
	time.Sleep(time.Second)
	fmt.Println("async result: ", asyncRes)
}

// 同步调用
func sync(methodName string, args *rpc_objects.Args, res *int, client *rpc.Client) {
	err := client.Call(methodName, args, res)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, res)
}

// 异步调用1
func async1(methodName string, args *rpc_objects.Args, res *int, client *rpc.Client) *rpc.Call {
	call := client.Go(methodName, args, res, nil)
	replyCall := <-call.Done
	return replyCall
}

// 异步调用2
func async2(methodName string, args *rpc_objects.Args, res *int, client *rpc.Client) {
	client.Go(methodName, args, res, nil)
}
