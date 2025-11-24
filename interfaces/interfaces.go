package main

import "fmt"

// interface implementation
type Animal interface {
	speak() string
}

// Dog struct
type Dog struct{}

func (d Dog) speak() string {
	return "woof!"
}

type Cat struct{}

func (c Cat) speak() string {
	return "Meow"
}

func MakeSpeak(a Animal) {
	fmt.Println(a.speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	MakeSpeak(dog)
	MakeSpeak(cat)
}
