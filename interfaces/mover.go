//implement a Mover interface with the mover() string function

package main

import (
	"fmt"
)

type Mover interface {
	Move() string
}

type Car struct {
	Brand string
}

func (c Car) Move() string {
	return "The car is moving"
}

type Person struct {
	Name string
}

func (p Person) Move() string {
	return "The person is walking"
}

func DescribeMovement(m Mover) {
	fmt.Println(m.Move())
}

func main () {
	car := Car{Brand: "Toyota"}
	person := Person{Name: "Lucky"}
	DescribeMovement(car)
	DescribeMovement(person)
}
