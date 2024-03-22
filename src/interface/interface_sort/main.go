package main

import (
	"fmt"
	"sort"
)

type StringArray []string

func (p StringArray) Len() int {
	return len(p)
}

func (p StringArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p StringArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {
	sa := StringArray{"1", "dasd", "fdskl", "0rew", "daad"}
	sort.Sort(sa)
	fmt.Println(sa)
}
