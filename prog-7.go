package main

import "fmt"

func main() {
	var (
		age  int
		name string
	)

	fmt.Println("Hello there!")
	fmt.Println("Welcome to Universe DB!\n")

	fmt.Printf("Please input your name: ")
	n, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(n, err)
	}
	fmt.Println("Your name is:", name)

	fmt.Printf("\nPlease input your age: ")
	n, err = fmt.Scanln(&age)
	if err != nil {
		fmt.Println(n, err)
	}

	for i := 0; i < age; i++ {
		fmt.Println("For the", i+1, "time, your name is:", name)
	}

	fmt.Println()

	i := 0
	for i < age {
		if i > 0 {
			break
		}
		fmt.Printf("Complementary print #%d\n", i+1)
		i++
	}

	numbers := [5]int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Printf("Range print #%d: %d\n", i+1, num)
	}

	for i, num := range [5]int{32, 54, 343, 2} {
		fmt.Printf("Range print (type 2) #%d: %d\n", i+1, num)
	}

	for i, num := range make([]int, 3) {
		fmt.Printf("Range print (type 3) #%d: %d\n", i+1, num)
	}

	fmt.Println()

}
