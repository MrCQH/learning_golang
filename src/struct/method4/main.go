package main

import (
	"fmt"
)

type Log struct {
	msg string
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}
func (l *Log) String() string {
	return l.msg
}

type Customer struct {
	Name string
	log  *Log
}

func (c *Customer) Log() *Log {
	return c.log
}

type Customer2 struct {
	Name string
	Log
}

func (c *Customer2) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log.String())
}

type Camera struct {
}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct {
}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	//combination()
	//embedded()
	multiEmbedded()
}

// 组合
func combination() {
	c := &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	// fmt.Println(c) &{Barak Obama 1 - Yes we can!}
	c.Log().Add("2 - After me the world will be a better place!")
	//fmt.Println(c.log)
	fmt.Println(c.Log())
}

// 继承
func embedded() {
	c := &Customer2{"Barak Obama", Log{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
}

// 多继承
func multiEmbedded() {
	pcp := new(CameraPhone)
	fmt.Println(pcp.Call(), pcp.TakeAPicture())
}
