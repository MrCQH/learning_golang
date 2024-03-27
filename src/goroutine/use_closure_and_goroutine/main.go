package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// 版本 A：
	for ix := range values { // ix 是索引值
		func() {
			fmt.Print(ix, " ")
		}() // 调用闭包打印每个索引值
	}
	fmt.Println()
	// 版本 B：和 A 版本类似，但是通过调用闭包作为一个协程
	// 因为这些闭包都只绑定到一个变量，这是一个比较好的方式，当你运行这段代码时，你将看见每次循环都打印最后一个索引值 4，而不是每个元素的索引值。
	// 因为协程可能在循环结束后还没有开始执行，而此时 ix 值是 4。
	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()
	time.Sleep(5e9)
	// 版本 C：正确的处理方式
	// 调用每个闭包时将 ix 作为参数传递给闭包。
	// ix 在每次循环时都被重新赋值，并将每个协程的 ix 放置在栈中，所以当协程最终被执行时，每个索引值对协程都是可用的。
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	fmt.Println()
	time.Sleep(5e9)
	// 版本 D：输出值：
	// 因为版本 D 中的变量声明是在循环体内部，所以在每次循环时，这些变量相互之间是不共享的，所以这些变量可以单独的被每个闭包使用。
	for ix := range values {
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	time.Sleep(1e9)
}
