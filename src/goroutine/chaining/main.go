package main

import (
	"flag"
	"fmt"
)

var nGoroutine = flag.Int("n", 100000, "how many goroutine")

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *nGoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	// 1. 该right是最终的right. 2. for 循环中最初的 go f(left, right) 因为没有发送者一直处于等待状态
	// 3. 当主线程的 right <- 0 执行时，类似于递归函数在最内层产生返回值一般
	right <- 0
	x := <-leftmost
	fmt.Println(x)
}
