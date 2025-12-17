package main


import (
	"io"
	"bytes"
	"os"
	"fmt"
)

func main() {
	fmt.Println("Type something and press enter: ")

	var memory bytes.Buffer

	io.Copy(&memory, os.Stdin)

	fmt.Println("You typed: ")
	fmt.Println(memory.String())
}

