package main


import (
	"io"
	"os"
	"strings"
)

func main() {
	text := strings.NewReader("Lucky is the best Golang software Engineer at Savannah Informatics\n")

	file, _ := os.Create("note.txt")
	defer file.Close()

	io.Copy(file, text)
}

