package main

import "fmt"

func main() {
	classifier(1, 2., "renzhen", 'c', true, nil)
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float32, float64:
			fmt.Printf("Param #%d is a float32 or float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		case byte:
			fmt.Printf("Param #%d is a byte\n", i)
		case rune:
			fmt.Printf("Param #%d is a rune\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
