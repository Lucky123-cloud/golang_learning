//reverse a slice without creating a new one

package main

import "fmt"

func ReverseSlice(nums []int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ReverseSlice(nums)
	fmt.Println(nums)
}

