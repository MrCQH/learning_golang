package main

import "fmt"

func main() {
	ch := make(chan int)
	defer close(ch)
	go func() {
		v, ok := <-ch
		if !ok {
			return
		}
		fmt.Println(v)
	}()
	ch <- 2
}
