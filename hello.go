package main

import (
	"fmt"
	"math/rand"
	"time"

	"rsc.io/quote"
)

func main() {
	var i = 9;
	x := 90;
	var f int;
	
	fmt.Println(i, x, f)

	fmt.Println("Hello, world!")
	fmt.Println(quote.Glass())

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
}
