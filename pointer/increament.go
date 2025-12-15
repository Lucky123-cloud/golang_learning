package main

import (
	"fmt"
)

func increament(x *int) {
	*x = *x + 1
}

func main() {
	a := 5
	increament(&a)
	fmt.Println(a)
}


