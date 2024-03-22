package main

import "fmt"

func main() {
	//str := "aaabbaag"
	//fmt.Println(str)
	//fmt.Println(string_reverse(str))
	//fmt.Println(uniq(str))
	//a := []int{1, 2, 3}
	res := _map(func(x int) int {
		return x * 10
	}, 1, 2, 3)
	fmt.Println(res)
}

func split_string(str string, i int) (string, string) {
	return str[:i+1], str[i+1:]
}

func string_split2(str string) string {
	return str[len(str)/2:] + str[:len(str)/2]
}

func string_reverse(str string) string {
	bytes := []byte(str)
	n := len(bytes)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-i-1] = bytes[n-i-1], bytes[i]
	}
	return string(bytes)
}

func uniq(str string) string {
	bytes := []byte(str)
	resBytes := make([]byte, 10)
	var last byte
	for _, iter := range bytes {
		if last == 0 {
			last = iter
			resBytes = append(resBytes, iter)
			continue
		}

		if last != iter {
			resBytes = append(resBytes, iter)
		}
		last = iter
	}
	return string(resBytes)
}

func _map(fn func(int) int, slice ...int) []int {
	var res []int
	for _, iter := range slice {
		res = append(res, fn(iter))
	}
	return res
}
