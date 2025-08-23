package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x + 1)
}

func main() {
	num := 5
	result := factorial(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
	testCount(1)
}
