package main

import (
	"fmt"
	"reflect"
)

func main() {
	//test1()
	test2()
}

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func test1() {
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	// alternative:
	// typ := value.Type()  // main.NotknownType
	fmt.Println(typ)
	knd := value.Kind() // struct
	fmt.Println(knd)
	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		// value.Field(i).SetString("C#")
	}
	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
}

type T struct {
	A int
	B string
}

func test2() {
	t := T{23, "skidoo"}
	v := reflect.ValueOf(&t).Elem()
	typeOfT := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// 字段
		f := v.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			// 变量名				类型			值
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	v.Field(0).SetInt(77)
	v.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
