package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("First one in the Stack")
	defer fmt.Println("Second one in the Stack")
	defer fmt.Println("Third one in the stack")

	fmt.Println("Hello World, we are starting here")
}
