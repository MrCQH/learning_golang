package main

import "fmt"

type myStruct struct {
	name   string
	age    int
	height float32
}

func main() {
	my1 := &myStruct{"mrchen", 24, 175}
	my2 := myStruct{"mrchen", 24, 175}
	fmt.Printf("%p %p\n", my1, &my2)
}
