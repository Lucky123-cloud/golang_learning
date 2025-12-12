//Reverse a String

package main


import (
	"fmt"
)

func ReverseMap(m map[string]int) map[int][]string{
	reversed := make(map[int][]string)

	for key, value := range m {
		reversed[value] = append(reversed[value], key)
	}
	return reversed
}

func main() {
	m := map[string]int{"apple":1, "banana":2, "avocado":1}
	out := ReverseMap(m)
	fmt.Println(out)
}
