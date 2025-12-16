package main

import (
	"fmt"
	"errors"
)

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	res, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	fmt.Println("The result is: ", res)
}

