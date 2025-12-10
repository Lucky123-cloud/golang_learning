//we can skip a condition if met and contniue, the loops does not stop


package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4, 5}
	for _, val := range num {
		if val != 3 {
			continue
		}
		fmt.Println(val)
	}
}

