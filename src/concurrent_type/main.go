package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name   string
	salary float64
	chF    chan func() // 用来代替控制并发 访问属性
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

// 通过通道去赋值Person的属性
// 只能够保证写与写之间的并发，不能保证写与读之间的并发
func (p *Person) SetSalary(sal float64) {
	p.chF <- func() {
		p.salary = sal
	}
}

func (p *Person) GetSalary() float64 {
	fCh := make(chan float64)
	p.chF <- func() {
		fCh <- p.salary
	}
	return <-fCh
}

func (p *Person) String() string {
	return "(" + p.Name + " " + strconv.FormatFloat(p.GetSalary(), 'f', 2, 64) + ")"
}

func main() {
	p1 := NewPerson("renzhen", 5555.55)
	fmt.Println(p1)
	p1.SetSalary(66666.66)
	fmt.Println(p1)
}
