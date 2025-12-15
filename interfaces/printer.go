//Create an interface Printer with:
//Print() string
//Write a function UsePrinter(p Printer) that prints the returned value.
//Pass two different structs

package main

import (
	"fmt"
)

type Printer interface {
	Print() string
}

type Invoice struct {
	ID int
}

func (i Invoice) Print() string {
	return fmt.Sprintf("Invoice ID: %d", i.ID)
}

type Report struct {
	MSG string
}

func (r Report) Print() string {
	return "Report Title: " + r.MSG
}

func UsePrinter(p Printer) {
	fmt.Println(p.Print())
}


func main() {
	invoice := Invoice{ID: 101}
	report := Report{MSG: "REPORTS FOR 2025"}

	UsePrinter(invoice)
	UsePrinter(report)
}

