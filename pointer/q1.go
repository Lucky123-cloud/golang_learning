//declare an integer variable x with value 50, then create a pointer p to x and print both the address and the value using the pointer

package main


import (
	"fmt"
)

func main() {
	var x int = 50
	var ptr *int = &x

	fmt.Println("Address of the pointer: ", &x)
	fmt.Println("Value of pointer: ", *ptr)
	fmt.Println("Value: ", ptr)
}

