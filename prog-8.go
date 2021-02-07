package main

import (
	"fmt"
	"math"
)

func main () {
	x, y := 3, 5

	// pass x by ref, y by val
	fmt.Println(addSub(&x, y))

	// function in function
	// lambda :: function as an rvalue
	var sqrt func(float64) float64 = func(num float64) float64 {
		return math.Sqrt(num)
	}

	fmt.Printf("Square root of %v is %.4f\n", x, sqrt(float64(x)))

	gen1 := generateDouble(4)
	fmt.Println(gen1())
	fmt.Println(gen1())
	fmt.Println(gen1())
}

// function
func addSub(a *int, b int) (sum, diff int) {
	sum = *a + b
	diff = *a - b

	return
}

func generateDouble(n int) func() int {
	num := n
	isFirstIt := true

	return func () int {
		if isFirstIt {
			isFirstIt = false
		} else {
			num *= 2
		}

		return num
	}
}
