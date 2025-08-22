package main

import (
	"fmt"
)

func main() {

	//multiple varibles declarations
	var a, b, c, d int = 1, 3, 5, 7
	var e, f = 6, "Hello"
	g, h := 7, "World!"

	//Group variables together
	var (
		i int
		j int    = 1
		k string = "Hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
