package main

import "fmt"

// var send_only chan<- int // channel can only receive data
// var recv_only <-chan int // channel can only send data

func main() {
	//test1()
	test2()
}

// 筛素数
func test2() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// 通道创建的时候都是双向的，但也可以分配给有方向的通道变量，就像以下代码：
func test1() {
	var c = make(chan int) // bidirectional
	go source(c)
	go sink(c)
}

func source(ch chan<- int) {
	for {
		ch <- 1
	}
}
func sink(ch <-chan int) {
	for {
		<-ch
	}
}
