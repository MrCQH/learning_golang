package main

import (
	_ "awesomeProject/lib1"
	_ "awesomeProject/lib2"
	"fmt"
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
