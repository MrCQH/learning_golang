package main

import "fmt"

type simple struct {
	a int
	b int
}

func main() {
	s := &simple{1, 2}
	fmt.Println(s.intAdd(3))
}
