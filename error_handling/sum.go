package main

import (
	"errors"
	"fmt"
)

var errorDivisionByZero = errors.New("You cannot divide by 0")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errorDivisionByZero
	}
	return a / b, nil
}

func main() {
	res, err := divide(10, 0)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return
	}
	fmt.Printf("result: %d\n", res)
}
