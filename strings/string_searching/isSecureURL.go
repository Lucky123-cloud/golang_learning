package main


import (
	"fmt"
	"strings"
)

//isSecureUrl returns true is the url string contains "https://" 
func isSecureURL(url string) bool {
	return strings.HasPrefix(url, "https://")
}

func main() {
	url := "https://example.com"
	fmt.Println(isSecureURL(url))
}


