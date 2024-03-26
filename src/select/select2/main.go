package main

import (
	"fmt"
)

func main() {
	//test1()
	test2()
}

func test2() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	go tel2(ch1, ch2)
	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case <-ch2:
			goto label
		}

	}
label:
	fmt.Println("End process!")
}

func tel2(ch1 chan int, ch2 chan bool) {
	for i := 0; i < 15; i++ {
		ch1 <- i
	}
	ch2 <- true
}

func test1() {
	ch := make(chan int)
	go tel(ch)
	for {
		// 尝试从空的channel里拿元素，会一直阻塞，导致死锁，触发panic
		// 需要close channel
		if v, ok := <-ch; ok {
			fmt.Println(v)
		}
	}
}

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch) // 尝试从空的channel里拿元素，会一直阻塞，导致死锁，触发panic, 需要close channel
}
