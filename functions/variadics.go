//return the sum of all the things


package main


import (
	"fmt"
)

func sum(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func main() {
	res := sum(3, 4, 5, 6, 7, 2, 3, 4, 2, 6, 7, 8)
	fmt.Println(res)
}

