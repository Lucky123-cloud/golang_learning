package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	//person 1
	pers1.name = "Shangwe"
	pers1.age = 10
	pers1.job = "Software Engineer"
	pers1.salary = 100000000

	//person 2
	pers2.name = "LUCKY"
	pers2.age = 12
	pers2.job = "CEO At Digibiashara"
	pers2.salary = 200000000

	//Access and print Person 1 details
	fmt.Println("Person 1 Details")
	fmt.Println("Name:", pers1.name)
	fmt.Println("Age:", pers1.age)
	fmt.Println("Job:", pers1.job)
	fmt.Println("Salary:", pers1.salary)
	fmt.Println()
	//Access and print Person 2 details
	fmt.Println("Person 2 Details")
	fmt.Println("Name:", pers2.name)
	fmt.Println("Age:", pers2.age)
	fmt.Println("Job:", pers2.job)
	fmt.Println("Salary:", pers2.salary)
	fmt.Println()
	// Print the entire struct
	fmt.Println("Person 1 Struct:", pers1)
	fmt.Println("Person 2 Struct:", pers2)
	fmt.Println()
	// Print the memory address of the struct
	fmt.Println("Person 1 Memory Address:", &pers1)
	fmt.Println("Person 2 Memory Address:", &pers2)
	fmt.Println()

	//Print Pers1 info by calling a function
	printPersonInfo(pers1)
	printPersonInfo(pers2)
}

func printPersonInfo(pers Person) {
	fmt.Println("Person Details from Function")
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
	fmt.Println()
}
