package main

import (
	"fmt"
	"testing"
)

func main() {
	//BenchmarkFunctionTest()
	BenchmarkSpecifiedCodeTest()
}

func BenchmarkSpecifiedCodeTest() {
	fmt.Println("Specified Code: ", testing.Benchmark(func(b *testing.B) {
		var cnt int
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			cnt++
		}
		b.StopTimer()
	}))
}

// 对函数做Benchmark测试
func BenchmarkFunctionTest() {
	fmt.Println("sync: ", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered: ", testing.Benchmark(BenchmarkChannelBuffered).String())
}

func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {

	}
}

func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	for range ch {

	}
}
