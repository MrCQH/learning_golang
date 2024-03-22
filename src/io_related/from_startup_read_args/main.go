package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	test1()
}

func test1() {
	who := "renzhen "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
	// runtime absPath 程序本身的名字
	fmt.Println("args[0]: ", os.Args[0])
}
