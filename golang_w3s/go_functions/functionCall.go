package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("Hello from myMessage function")
}

func main() {
	myMessage() //function call
	myMessage()
	myMessage()
}
