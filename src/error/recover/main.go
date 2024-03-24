package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	test2()
}

func test2() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	// 不会执行了，执行完defer后，跳出当前函数，继续执行
	fmt.Printf("After bad call\r\n") // <-- would not reach
}

func badCall() {
	panic("bad end")
}

func test1() {
	protect(func() error {
		panic("你好，panic")
		return errors.New("你好, error")
	})
	test1()
}

func protect(g func() error) {
	defer func() {
		log.Println("done")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			// this is panic, not error
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g() //   possible runtime-error
}
