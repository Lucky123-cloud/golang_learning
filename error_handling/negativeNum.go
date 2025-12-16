package main

import (
	"fmt"
	"errors"
)

var ErrNegativeNum = errors.New("The number is a negative number")

func checkNumber(n int) error {
	if n < 0 {
		return ErrNegativeNum
	}
	return nil
}

func main() {
	err := checkNumber(-5)
	if err == ErrNegativeNum {
		fmt.Println("Error occured: ", err)
	}
}
