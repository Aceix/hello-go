package main

import (
	"fmt"
	"strings"
)

func main() {
	name := ""

	fmt.Printf("Enter your first name: ")
	fmt.Scanf("%s", &name)

	fmt.Println("Your name has", len(name), "characters")
	fmt.Printf("Your new name is %s\n", strings.Join([]string{name, "Smith"}, " "))

	fmt.Printf("Average = %.2f\n", getAverage([]int{1000, 2, 3, 17, 50}))

	print3LenIntArr([3]int{4, 23})
}

func getAverage(nums []int) (avg float32) {
	size := len(nums)
	var acc = 0

	for i := 0; i < size; i++ {
		acc += nums[i]
	}

	avg = float32(acc) / float32(size)

	return
}

func print3LenIntArr(nums [3]int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Value at index %d: %d\n", i, nums[i])
	}
}
