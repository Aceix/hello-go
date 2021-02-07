package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y int
	r    float32
}

func (c Circle) area(newRadius float32) float32 {
	if newRadius != 0 {
		return math.Pi * newRadius * newRadius
	}
	return math.Pi * c.r * c.r
}

var a int
// B is a global variable
var B = 90

func main() {
	c1 := Circle{
		r: 5,
	}
	c2 := Circle{0, 0, 4}

	fmt.Println("Area of c1: ", c1.area(0))
	fmt.Println("Also adjusted area of c1: ", Circle.area(c1, 15))
	fmt.Println("Area of c2: ", c2.area(0))
}
