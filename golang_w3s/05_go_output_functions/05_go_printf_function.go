package main

// The Printf() function first formats its argument based on the given formatting verb and then prints them.

// Verb	Description
// %v	Prints the value in the default format
// %#v	Prints the value in Go-syntax format
// %T	Prints the type of the value
// %%	Prints the % sign

import (
	"fmt"
)

func main() {
	// %v is used to print the value of the arguments
	// %T is used to print the type of the arguments
	var i string = "Hello"
	var j int = 15
	var k = 15.5
	var text = "Hello World!"

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T\n", j, j)
	fmt.Printf("%v\n", k)
	fmt.Printf("%#v\n", k)
	fmt.Printf("%v%%\n", k)
	fmt.Printf("%T\n", k)

	fmt.Printf("%v\n", text)
	fmt.Printf("%#v\n", text)
	fmt.Printf("%T\n", text)
}
