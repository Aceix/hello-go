package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Local().UnixNano())
	
	mark := rand.Intn(101)
	var grade string

	fmt.Println("Mark:", mark)

	// "break" is not needed in golang
	// the program automatically breaks as expected

	// also, defualt switch expr is "true"
	// ie: the line below is same as
	// writing "switch(true)""
	switch {
	case mark >= 90:
		grade = "A+"
		// break
	case mark >= 70:
		grade = "A"
	case mark >= 60:
		grade = "B"
	case mark >= 50:
		grade = "C"
	case mark >= 45:
		grade = "D"
	case mark >= 40:
		grade = "E"
	default:
		grade = "F"
	}

	fmt.Println("Final grade:", grade)

}
