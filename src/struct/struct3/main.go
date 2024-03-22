package main

import (
	"fmt"
)

type Birth struct {
	month int
	day   int
	year  int
}

type VCard struct {
	name     string
	addr     string
	birth    Birth
	imageUrl string
}

func main() {
	pCard := &VCard{name: "mrchen", addr: "chengdou", birth: Birth{6, 6, 1999}, imageUrl: "www.example.com"}
	fmt.Println(pCard)
}
