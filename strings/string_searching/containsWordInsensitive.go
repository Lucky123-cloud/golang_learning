package main

import (
	"fmt"
	"strings"
)

func containsWordCaseInsensitive(s, word string) bool {
	return strings.Contains(
		strings.ToLower(s),
		strings.ToLower(word),
	)
}
func main() {
	str := "Hello Go is really a great Language"
	word := "go"
	fmt.Println(containsWordCaseInsensitive(str, word))
}

