package main

import (
	"fmt"
	"time"
)

func main() {
	test2()
}

// 简单的定时任务，方式一
func test1() {
	ticker := time.NewTicker(1e9)
	defer ticker.Stop()
	for {
		fmt.Println(<-ticker.C)
	}
}

// 方式二
func test2() {
	tick := time.Tick(time.Second)
	for {
		fmt.Println(<-tick)
	}
}
