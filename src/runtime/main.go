package main

import (
	"fmt"
	"runtime"
)

type test struct {
	a int
	b string
}

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
	t := &test{1, "1"}

	runtime.SetFinalizer(t, func(t *test) {
		fmt.Println("final", t)
	})
	runtime.GC()
}
