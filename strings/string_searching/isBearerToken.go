package main


import (
	"fmt"
	"strings"
)

func isBearerToken(token string) bool {
	return strings.HasPrefix(token, "Bearer")
}

func main() {
	token := "Bearer djfijbfjierwibisibhfshb"
	fmt.Println(isBearerToken(token))
}

