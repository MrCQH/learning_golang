package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

func (this *Car) numberOfWheels() int {
	return this.wheelCount
}

type Mercedes struct {
	Car
}

func (this *Mercedes) Start() {
	fmt.Printf("begin start...")
}

func (this *Mercedes) Stop() {
	fmt.Printf("begin start...")
}

func (this *Mercedes) sayHiToMerkel() {
	fmt.Printf("Hi Mercedes: %v", this)
}

func main() {
	m := &Mercedes{Car: Car{nil, 4}}
	m.sayHiToMerkel()
	m.Start()
}
