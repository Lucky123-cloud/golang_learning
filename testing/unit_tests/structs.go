package main

import (
	"fmt"
)

type Company struct {
	Name             string
	YearOfFoundation string
	NoOfEmployees    int
	NoOfDirectors    int
	NoOfDepartments  int
	Departments      []string
	NoOfComputers    int
}

func main() {
	var comp Company
	// var comp2 Company

	comp.Name = "Digibiashara"
	comp.YearOfFoundation = "Jan 1, 2025"
	comp.NoOfEmployees = 10
	comp.NoOfDirectors = 4
	comp.Departments = []string{"IT", "Finance", "SWE", "Security", "Accomodation"}
	comp.NoOfDepartments = 5
	comp.NoOfComputers = 300

	// print comp info
	fmt.Println(comp.Name)

	//print from a function
	PrintFromAFunction(comp)

}

func PrintFromAFunction(company Company) {
	fmt.Println("Printing Company Info from a Function")
	fmt.Println("Name of Company: ", company.Name)
	fmt.Println("Year Of Foundation: ", company.YearOfFoundation)
	fmt.Println("Number of Employees: ", company.NoOfEmployees)
	fmt.Println("Number of Directors: ", company.NoOfDirectors)
	fmt.Println("Number of Departments: ", company.NoOfDepartments)
	fmt.Println("Company departments in structs: ", company.Departments)
	fmt.Println("Number of Computers: ", company.NoOfComputers)
}
