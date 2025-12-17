package main


import (
	"fmt"
	"strings"
)

func isJSONfile(fileName string) bool {
	return strings.HasSuffix(fileName, ".json")
}

func main() {
	fileName := "logs.json"
	fmt.Println(isJSONfile(fileName))
}


