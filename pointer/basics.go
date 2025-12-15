//basics of pointers
//pointer - Is a variable that stores the memory address of another variable instead of the value itself

package main

import (
	"fmt"
)

func main() {
	var num int = 5 //the num variable holds the value 5
	var p *int //the variable p holds a memory address of certain integer
	var ptr *int
	p = &num
	*p = 20

	fmt.Println("Address of num: ", &p)
	fmt.Println("Pointer to num: ", p)
	fmt.Println("Value of num: ", *p)
	fmt.Println("Nil Pointer: ", ptr)
}


