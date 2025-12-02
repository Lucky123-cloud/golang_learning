package main

import (
	"fmt"
	"errors"
)

type NegativeNumberError struct {
	Number int
	Message string
}

func (e *NegativeNumberError) Error() string {
	return fmt.Sprintf("Negative number detected %d (%s)", e.Number, e.Message)
}

//checkNegNum checks whether a number is negative or positive and returns string
func checkNegNum(n int) (string, error) {
	if n < 0 {
		//wrap custom error
		return "", fmt.Errorf("Error occured: %w", &NegativeNumberError {
			Number: n,
			Message: "The number is a negative number",
		})
	}
	return "The number is a positive number", nil
}

func main() {
	res, err := checkNegNum(-5)
	if err != nil {
		var negError *NegativeNumberError
		if errors.As(err, &negError) {
			fmt.Println("We caught a negative number Error")
			fmt.Println("Number: ", negError.Number)
			fmt.Println("Message: ", negError.Message)
			return
		}
		fmt.Println("Unknown Error Occured, ", err)
		return
	}
	fmt.Println("Hello the result is: ", res)
}
