//return a new slice with only even numbers


package main


import (
	"fmt"
)

func FilterEven(nums []int) []int {
	evensOnly := []int{}
	for _, val := range nums {
		if val%2 == 0 {
			evensOnly = append(evensOnly, val)
		}
	}
	return evensOnly
}

func main() {
	nums := []int{2, 3, 4, 6, 7, 8, 9, 45, 32, 46}
	res := FilterEven(nums)
	fmt.Println(res)
}

