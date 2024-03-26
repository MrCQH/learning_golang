package main

import (
	"fmt"
	"strconv"
)

type Request struct {
	a, b   int
	replyc chan int
}

func (r Request) String() string {
	return "(" + strconv.Itoa(r.a) + " " + strconv.Itoa(r.b) + ")"
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, server chan *Request, quit chan bool) {
	for {
		select {
		case req := <-server:
			go run(op, req)
		case <-quit:
			return
		}
	}
}

func startServer(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)
	go server(op, service, quit)
	return service, quit
}

func main() {
	serviceCh, quitCh := startServer(func(a, b int) int {
		return a + b
	})
	// 构造请求 100个请求
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		serviceCh <- req
	}

	for i := N - 1; i >= 0; i-- {
		if v := <-reqs[i].replyc; v != N+i+i {
			fmt.Println("fail at ", reqs[i])
		} else {
			//fmt.Println("Request ", reqs[i], " is ok!")
			fmt.Printf("%v %v\n", reqs[i], v)
		}
	}
	quitCh <- true
	fmt.Println("done")

}
