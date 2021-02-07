package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}

func split(sum int) (x, y int) {
	x = sum / 4
	y = sum * 3 / 4
	return
}

// package-level variables
var t1, t2, t3 bool

func main() {
	// function variables
	var x = 9
	var y = 49

	fmt.Println("Featured addition")
	fmt.Print(x, " + ", y, " = ")
	fmt.Println(add(x, y))

	fmt.Println("Featured subtraction")
	fmt.Print(x, " - ", y, " = ")
	fmt.Println(subtract(x, y))

	var a = "hello"
	var b = "world"

	fmt.Println("Swapping %s and %s", a, b)
	fmt.Println(swap(a, b))

	c, d := swap(a, b)
	fmt.Println("Create assignment %s and %s", c, d)

	fmt.Println("Using naked return in functions:")
	fmt.Println(split(30))

	// multi var initialization
	var i, j int = 1, 2              // single type
	var t1, t2, t3 = "ama", 2, false // multi type

	fmt.Println(i, j)
	fmt.Println(t1, t2, t3)

	// short var declaration
	k := 90
	m, n, o := "Ama", "Kwabena", "Esi"

	fmt.Println(k)
	fmt.Println(m, n, o)

}
