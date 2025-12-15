//Write a function doubleVal that doubles the value of an integer using a pointer

package main

import (
	"fmt"
)

func doubleVal(x *int) {
	*x = *x * 2
}

func main() {
	a := 50
	doubleVal(&a)
	fmt.Println(a)
}

