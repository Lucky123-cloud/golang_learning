//we could iterate string to get individual characters


package main

import (
	"fmt"
)

func main () {
	str := "Lucky"

	for i, ch := range str {
		fmt.Println(i, string(ch))
	}
}
