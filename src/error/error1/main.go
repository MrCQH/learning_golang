package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	if sqrt, err := Sqrt(-3.4); err != nil {
		log.Fatal("Error: ", err)
	} else {
		log.Println(sqrt)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// error 不要以大写的英语字母开头
		return 0, fmt.Errorf("math - square root of negative number %g", f)
	}
	return math.Sqrt(f), nil
}
