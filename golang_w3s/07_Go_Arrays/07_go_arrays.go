package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	var arr3 = [...]int{1, 2, 3}
	arr := [...]int{4, 5, 6, 7, 8}

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	prices := [3]int{10, 20, 30}
	prices[2] = 50

	arr4 := [5]int{}              //not initialized
	arr5 := [5]int{1, 2}          //partially initialized
	arr6 := [5]int{1, 2, 3, 4, 5} //fully initilized
	arr7 := [5]int{1: 10, 2: 40}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr)
	fmt.Println(cars)
	fmt.Println(prices[0])
	fmt.Println(prices[1])
	fmt.Println(prices[2])
	fmt.Println(prices)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)
	fmt.Println(len(arr1))
	fmt.Println(len(arr))
}
