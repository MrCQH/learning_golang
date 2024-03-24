package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func init() {
	check()
}

func main() {
	//test1()
}

func check() {
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
}

func test1() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}
