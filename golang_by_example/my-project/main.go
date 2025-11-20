package main

import (
	"fmt"
	"my-project/mathematic/ingo"
)

func main() {
	sum := ingo.Add(4, 6)
	product := ingo.Multiply(4, 6)

	fmt.Println(sum)
	fmt.Println(product)

	diff := ingo.Subtract(3, 4)
	fmt.Println(diff)

	fmt.Println(ingo.Version)

	a, b := ingo.Swap("Baraka", "Lucky")
	fmt.Println(a, b)

	fmt.Println(ingo.Split(25))
}
