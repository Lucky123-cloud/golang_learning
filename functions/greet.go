//write a function that returns Hello, <name>

package main

import (
	"fmt"
)


func Greet(name string) string{
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	res := Greet("lucky")
	fmt.Println(res)
}
