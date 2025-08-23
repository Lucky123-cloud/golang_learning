package main

import (
	"fmt"
)

func main() {
	//Making slices
	//1. Using th syntax: => var myslice []datatype
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	//creating slices from arrays
	// myslice := myarray[start:end]

	myarray := [6]int{10, 20, 30, 40, 50, 60}
	myslice3 := myarray[2:4]

	fmt.Printf("myslice = %v\n", myslice3)
	fmt.Printf("Length = %d\n", len(myslice3))
	fmt.Printf("Capacity = %d\n", cap(myslice3))

	//2. Using the make function
	// make([]datatype, length, capacity)
	myslice4 := make([]int, 5, 10)
	fmt.Printf("Myslice4 = %v\n", myslice4)
	fmt.Printf("Length = %d\n", len(myslice4))
	fmt.Printf("Capacity = %d\n", cap(myslice4))
	fmt.Println(myslice4)

	//with omitting capacity
	myslice5 := make([]int, 5)
	fmt.Printf("Myslice5 = %v\n", myslice5)
	fmt.Printf("Length = %d\n", len(myslice5))
	fmt.Printf("Capacity = %d\n", cap(myslice5))
	fmt.Println(myslice5)
}
