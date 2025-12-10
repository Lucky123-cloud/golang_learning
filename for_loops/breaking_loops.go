//when we break loops we simply want to break out when a certain condition is met - the loop does not contniue

package main


import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4, 5}
	for _, val := range num {
		if val%2 == 0 {
			break
		}
		fmt.Println(val)
	}
}
//it stops immediately the condition is met

