package main

import (
	"fmt"
	_ "learning_golang/lib1"
	_ "learning_golang/lib2"
)

func main() {
	a, b, c := 1, 2, 3
	arr1 := []*int{&a, &b, &c}
	arr2 := make([]*int, 3)
	copy(arr2, arr1)
	fmt.Printf("arr1 addr: %p, arr2 addr: %p\n", arr1[0], arr2[0])
	fmt.Printf("arr1 value: %v, arr2 value: %v\n", *arr1[0], *arr2[0])
}

func init() {
	println("main init...")
}
