package main

import (
	"fmt"
)


func main() {
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println("Numbers: ", num)
	}
}

//we can do the same when we only want the value alone and nothing else
