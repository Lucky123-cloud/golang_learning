package main

import (
	"fmt"
	"strings"
)

//ContainsGo takes in a string and returns a true if the string contains go
func ContainsGo(s string) bool {
	return strings.Contains(s, "Go")
}

func main() {
	s := "Go is a great language"
	fmt.Println(ContainsGo(s))
}
