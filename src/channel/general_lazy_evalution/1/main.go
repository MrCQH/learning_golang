package main

import (
	"fmt"
)

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}
func generateInteger() int {
	return <-resume
}

// 无前值依赖的generator
func main() {
	resume = integers()
	for i := 0; i < 100; i++ {
		fmt.Println(generateInteger())
	}
	//fmt.Println(generateInteger()) //=> 0
	//fmt.Println(generateInteger()) //=> 1
	//fmt.Println(generateInteger()) //=> 2
}
