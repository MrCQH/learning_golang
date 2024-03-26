package main

import (
	"fmt"
	"time"
)

func main() {
	//test1()
	test2()
}

func test2() {
	suck(pump())
	time.Sleep(1e9)
}

// 返回一个只读的通道
func pump() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 返回一个只读的通道
func suck(ch <-chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

// 信号量模式: 模型概要
func signalModel1() {
	ch := make(chan int)
	go func() {
		// doSomething
		ch <- 1 // Send a signal; value does not matter
	}()
	//doSomethingElseForAWhile()
	<-ch // Wait for goroutine to finish; discard sent value.
}

// 阻塞，该channel为无缓冲channel
func test1() {
	out := make(chan int)
	out <- 2
	go func(in chan int) {
		fmt.Println(<-in)
	}(out)
	time.Sleep(1 * time.Second)
}
