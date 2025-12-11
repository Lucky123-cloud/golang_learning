//Given an array of  integers, return the largest number


package main

import (
	"fmt"
)

func MaxArray(arr [6]int) int {
	maxi := arr[0]
	for i := 1; i < len(arr); i++{
		if arr[i] > maxi {
			maxi = arr[i]
		}
	}
	return maxi
}

func main() {
	arr := [6]int{1, 5, 6, 4, 7, 8}
	res := MaxArray(arr)
	fmt.Println(res)
}
