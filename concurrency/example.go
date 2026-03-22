package main

import (
	"fmt"
)

func main() {
	type rune = int32
	type byte = uint8

	i := 42
	r := byte(i)

	fmt.Println(r)
}
