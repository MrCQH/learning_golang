package main

import (
	"fmt"
	"time"
)

const MaxReqs = 50

var sem = make(chan int, MaxReqs)

type Request struct {
	a, b    int
	replych chan int
}

func server(serviceCh chan *Request, quitCh chan bool) {
	for {
		select {
		case v := <-serviceCh:
			go handle(v)
		case <-quitCh:
			return
		}
	}
}

func handle(req *Request) {
	sem <- 1
	go process(req)
	<-sem
	time.Sleep(time.Second)
}

func process(req *Request) {
	req.replych <- req.a + req.b
}

func main() {
	serviceCh := make(chan *Request)
	quitCh := make(chan bool)
	go server(serviceCh, quitCh)

	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replych = make(chan int)
		serviceCh <- req
	}

	for i := 0; i < N; i++ {
		if res, ok := <-reqs[i].replych; ok {
			fmt.Println(res)
		}
	}

	quitCh <- true
	fmt.Println("done")
}
