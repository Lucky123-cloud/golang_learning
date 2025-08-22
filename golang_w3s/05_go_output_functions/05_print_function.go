//prints arguments with their default format

package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World!"
	var k, l = 10, 20

	fmt.Print(i)
	fmt.Print(j)
	//result HelloWorld!
	fmt.Print("\n")

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
	// result: Hello
	// 		   World!

	fmt.Print(i, "\n", j)
	// result: Hello
	// 		   World!
	fmt.Print("\n")

	fmt.Print(i, " ", j)
	//result: Hello World!
	fmt.Print("\n")

	fmt.Print(k, l)
	//result: 10 20 Print() inserts a space between the arguments if neither are strings:

}
