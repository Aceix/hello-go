package main

import (
	"fmt"
	"log"
)

func main() {
	var num int

	fmt.Print("Please enter an integer: ")

	if _, err := fmt.Scanf("%d", &num); err != nil {
		log.Fatalln("An error occured:", err)
	}

	fmt.Println()

	switch {
	case num < 0:
		fmt.Println("Number is negative")
	case num == 0:
		fmt.Println("Number is zero")
	case num > 0:
		fmt.Println("Number is positive")
		fallthrough
	case num < 50:
		fmt.Println("Number could be less than 50. Maybe more.. hehe")
	}
}
