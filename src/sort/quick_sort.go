package main

import "fmt"

func main() {
	a := make([]int, 5)
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}

	qSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func qSort(a []int, l int, r int) {
	if l >= r {
		return
	}
	mid := a[(l+r)>>1]
	i, j := l-1, r+1
	for i < j {
		for i += 1; a[i] < mid; {
			i++
		}
		for j -= 1; a[j] > mid; {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	qSort(a, l, j)
	qSort(a, j+1, r)
}
