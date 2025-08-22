package main

import (
	"fmt"
)

//The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:

func main() {
	var i, j string = "Hello", "World"
	var k string = "Hello"
	var l int = 15

	// %v is used to print the value of the arguments
	// %T is used to print the type of the arguments

	fmt.Print("i has value: %V and type: %T\n", k, k)
	fmt.Print("l has value: %V and type: %T\n", l, l)

	fmt.Println(i, j)
}
