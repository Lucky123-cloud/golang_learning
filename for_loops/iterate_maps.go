//we can iterate maps, and for this, we need to know that we have keys and values, so we are trying to get the both of them

package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"a":1, "b":2, "c":3, "d":4, "e":5}
	for key, value := range m {
		fmt.Println("Key: ", key, "value: ", value)
	}
}
