//for loop with range basically helps you loop over a collection, like maps, arrays/slices, strings etc

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		fmt.Println("index: ", i, "Number: ", num)
	}
}
