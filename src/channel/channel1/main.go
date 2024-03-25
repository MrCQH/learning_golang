package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go getData(ch)
	go sendData(ch)
	time.Sleep(10 * time.Second)
}

func getData(ch chan string) {
	var input string
	time.Sleep(2 * time.Second)
	for {
		input = <-ch
		fmt.Printf("channel's value is: %s\n", input)
	}
}

func sendData(ch chan string) {
	ch <- "mrchen"
	ch <- "renzhen"
	ch <- "a"
	ch <- "b"
	ch <- "c"
}
