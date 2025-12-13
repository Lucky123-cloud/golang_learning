//anonymous func

package main


import (
	"fmt"
)


func main() {
	square := func(n int) int {
		return n * n
	}(7)
	fmt.Println(square)
}

