package main

import (
	"fmt"
	"math"
)

type AbsInterface interface {
	Abs() int
}

type SqrInterface interface {
	Sqr() float64
}

type Point struct {
	x int
}

func (p *Point) Abs() int {
	if p.x < 0 {
		return -p.x
	}
	return p.x
}

func (p *Point) Sqr() float64 {
	return math.Sqrt(float64(p.x))
}

func main() {
	pp := new(Point)
	pp.x = 4
	var ai AbsInterface
	var si SqrInterface
	var empty any
	empty = pp
	ai = empty.(AbsInterface)
	si = empty.(SqrInterface)
	empty = si

	fmt.Printf("ai pointer address: %p, abs is %d\n", ai, ai.Abs())
	fmt.Printf("si pointer address: %p, sqr is %.2f\n", si, si.Sqr())
}
