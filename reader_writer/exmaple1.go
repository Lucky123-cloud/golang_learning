package main

import (
	"os"
	"io"
)

func main() {
	file, _ := os.Open("hello.txt")
	defer file.Close()

	io.Copy(os.Stdout, file)
}

