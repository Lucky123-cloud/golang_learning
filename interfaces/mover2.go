package main

import (
	"fmt"
)

type Mover interface{
	Move() string
}

type Person struct {
	Name string
}

func (p Person) Move() string {
	return "Hello there, I am walking"
}

type Car struct {
	Brand string
}

func (c Car) Move() string {
	return "Hello there, the car is driving"
}

func DecideMovement(m Mover) {
	fmt.Println(m.Move())
}

func main() {
	car := Car{Brand: "Toyota"}
	person := Person{Name: "Lucky"}

	DecideMovement(car)
	DecideMovement(person)
}
