package main

import "fmt"

//Go requires explicit returns
func plus(a int, b int) int {
	return a + b
}

//When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func testfunctions() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}
