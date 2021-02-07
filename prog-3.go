package main

import "fmt"

func main() {
	var a int32 = 9
	var b, c, d = 1, 4, 6
	var e float32 = 90.3
	var f complex128 = 9 - 4.2i
	var g byte = 12
	var h rune = 90
	var i uintptr = 5
	var j bool
	k := 32.8
	const L int = 328
	const M string = "hello world"

	const (
		N        = 23
		O string = "CONSTANT"
	)

	fmt.Println("sum =", int64(a)+int64(b)+int64(c)+int64(d))
	fmt.Println("float e:", e)
	fmt.Println("complex f:", f)
	fmt.Println("byte g:", g)
	fmt.Println("rune h:", h)
	fmt.Println("uintptr i:", i)
	fmt.Println("bool j:", j)
	fmt.Printf("%T k: %.2f\n", k, k)
	fmt.Printf("%T l: %d\n", L, L)
	fmt.Printf("%T m: %s\n", M, M)
	fmt.Printf("%T N: %d\n", N, N)
	fmt.Printf("%T O: %s\n", O, O)

	// lvalue is data
	// rvalue is a memory lcation
}
