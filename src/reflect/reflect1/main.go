package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	// 一个结构体
	var _type reflect.Type = reflect.TypeOf(x)
	var _value reflect.Value = reflect.ValueOf(x)
	fmt.Println(_type, _value)

	// 定义的枚举
	fmt.Println(_type.Kind())
	fmt.Println(reflect.Float64 == _type.Kind())
	fmt.Println(_value.Kind() == _type.Kind())
	fmt.Println(reflect.Float64 == 14)
	fmt.Println(_type.Kind() == 14)

	// value
	fmt.Println("float value is : ", _value.Float())
	fmt.Println("float type is : ", _value.Kind())

	// 得到接口的值
	fmt.Println(_value.Interface())
	fmt.Println("--------")

	test()
}

func test() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
