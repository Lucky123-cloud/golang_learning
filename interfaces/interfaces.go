package main

import (
	"fmt"
)

// interface implementation
type Speaker interface {
	speak() string
}

// Dog struct
type Dog struct{}

func (d Dog) speak() string {
	return "woof!"
}

type Cat struct{}

func (c Cat) speak() string {
	return "Meow!"
}

func main() {
	var s Speaker

	s = Dog{}
	fmt.Println(s.speak())

	s = Cat{}
	fmt.Println(s.speak())
}
