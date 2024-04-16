package main

import "fmt"

type Value interface {
	int | int32 | float32 | float64
}

type MySlice[T Value] []T

func main() {
	m := MySlice[float64]{1.3, 2, 3}
	fmt.Println(m)
}
