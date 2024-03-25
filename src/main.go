package main

import (
	"fmt"
	_ "learning_golang/lib1"
	_ "learning_golang/lib2"
	"strconv"
)

func main() {
	var input string
	n, _ := fmt.Scan(&input)
	if n > 0 {
		convInput, _ := strconv.Atoi(input)
		println(convInput)
	}
}

func init() {
	println("main init...")
}

func compute(a, b int) (int, int) {
	return a + b, a - b
}
