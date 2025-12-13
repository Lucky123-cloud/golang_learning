//Write a function to compute the area of a rectangle
package main

import (
	"fmt"
)
//Area: computes the area of a rectangle
func Area(width, height float64) float64 {
	return width * height
}

func main() {
	res := Area(5, 10)
	fmt.Println(res)
}

