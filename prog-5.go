package main

import "fmt"

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("Type of i is %T", i)
	case int:
		fmt.Printf("Type of i is %T", i)
	case int32:
		fmt.Printf("Type of i is %T", i)
	case bool, string:
		fmt.Printf("Type of i is %T", i)
	case func(int) float32:
		fmt.Printf("Type of i is %T", i)
	}
}
