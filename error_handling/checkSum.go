package main

import (
	"fmt"
	"errors"
)

func checkNum(n int) (string, error) {
	if n < 0 {
		return "", errors.New("Your number is a negative number")
	}
	return "Your number is a positive number", nil
}

func main() {
	res, err := checkNum(5)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}
	fmt.Println("Your result is: ", res)
}

