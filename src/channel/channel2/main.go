package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	cnt := 0
	go prod(ch, &cnt)
	go inc(&cnt)
	go consum(ch)
	time.Sleep(30 * time.Second)
}

func consum(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func inc(cnt *int) {
	for {
		time.Sleep(time.Second)
		*cnt++
		fmt.Println("cur time: ", *cnt)
	}
}

func prod(ch chan int, cnt *int) {
	for i := 0; ; i++ {
		if *cnt == 15 {
			break
		}
		time.Sleep(time.Second)
		ch <- i
	}
}
