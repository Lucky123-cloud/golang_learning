package main

import (
	"fmt"
)

func SafeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("You cannot divide by 0")
	}
	return a / b, nil
}

func main() {
	res, err := SafeDivide(8, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(res)
}
