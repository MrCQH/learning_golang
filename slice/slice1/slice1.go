package main

func main() {

}

func magnify_slice(s []int, factor int) []int {
	n := len(s)
	nn := n * factor
	ns := make([]int, nn)
	copy(ns, s)
	return ns
}

func filter(s []int, fn func(int) bool) []int {
	var ns []int
	for _, iter := range s {
		if fn(iter) {
			ns = append(ns, iter)
		}
	}
	return ns
}

func RemoveStringSlice(s []int, start, end int) []int {
	ns := s[0:start]
	return ns
}
