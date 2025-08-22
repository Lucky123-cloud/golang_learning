package main

import (
	"fmt"
)

const PI = 3.14
const A int = 1 //typed constant
const B = 2     //untyped constant

const (
	C int    = 3
	D        = 4
	E string = "Hello"
	F        = "World!"
)

func main() {
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
