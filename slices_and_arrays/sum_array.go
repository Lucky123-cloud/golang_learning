//write a function that takes an array of 5 integers and returns the sum

package main


import (
	"fmt"
)


func sumArray(arr [5]int) int {
	sum := 0
	for i := 0; i < len(arr); i++{
		sum += arr[i]
	}
	return sum
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	res := sumArray(arr)
	fmt.Println(res)
}


	
