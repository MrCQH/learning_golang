package main

import "fmt"

type Rectangle struct {
	high  int
	width int
}

func (this *Rectangle) Area() int {
	return this.width * this.high
}

func (this *Rectangle) Perimeter() int {
	return 2 * (this.width + this.high)
}

func main() {
	rectangle := &Rectangle{2, 2}
	fmt.Println(rectangle.Area(), rectangle.Perimeter())

}
