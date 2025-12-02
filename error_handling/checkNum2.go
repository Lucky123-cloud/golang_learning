package main

import (
	"fmt"
	"errors"
)


var negativeNum = errors.New("The number is a negative number")

func checkNum(n int) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("There is probably something wrong with %d, lets see trace here: %w", n, negativeNum)
	}
	return "Your number is actually a positive number", nil
}

func main() {
	res, err := checkNum(-8)
	if err != nil {
		if errors.Is(err, negativeNum) {
			fmt.Println("Something is probably wrong, let's examine", err)
			rootError := errors.Unwrap(err)
			fmt.Println("The root cause is: ", rootError)
			return
		}
		fmt.Println("This is an unknown error", err)
		return 
	}
	fmt.Println("The result is: ", res)
}

