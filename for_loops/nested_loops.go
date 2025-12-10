package main


import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		for j := 1; j < 3; j++ {
			fmt.Println("i: ", i, "j: ", j)
		}
	}
}
