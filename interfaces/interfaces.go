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

// lets now learn how to accept interfaces to a function
func MakeSpeak(s Speaker) {
	fmt.Println(s.speak())
}

func main() {
	var s Speaker

	s = Dog{}
	fmt.Println(s.speak())
	MakeSpeak(Cat{})

	s = Cat{}
	fmt.Println(s.speak())
	MakeSpeak(Dog{})

}
