package main

import (
	"fmt"
	"errors"
)

func checkEven(n int) error {
	if n%2 != 0 {
		return errors.New("This is an odd number")
	}
	return nil
}

func main() {
	err := checkEven(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number is even")
	}
}
