//return the sum of the elems in the slice

package main


import (
	"fmt"
)

func sumSlice(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func main() {
	nums := []int{1, 232, 4, 5, 6}
	res := sumSlice(nums)
	fmt.Println(res)
}

