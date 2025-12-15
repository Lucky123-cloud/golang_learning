package main

import (
	"fmt"
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x, y := 5, 10
	fmt.Println("Values before swapping: ", x, y)

	swap(&x, &y)

	fmt.Println("values after swapping: ", x, y)
}
