package main

import (
	"fmt"
	"errors"
)

type NotFoundError struct {
	Resource string
	ID int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Resource %s with ID %d not found ", e.Resource, e.ID)
}

func FindResourceWithID(id int) (string, error) {
	if id == 0 {
		return "", fmt.Errorf("Error Occured: %w", &NotFoundError{
			Resource: "User",
			ID: id,
		})

	}
	return "Resource found", nil
}

func main() {
	res, err := FindResourceWithID(-1)
	if err != nil {
		var nfErr *NotFoundError
		if errors.As(err, &nfErr) {
			fmt.Println("We caught an error")
			fmt.Println("We did not found resource with ID: ", nfErr.ID)
			return
		}
		fmt.Println("This error is not known: ", err)
		return	
	}
	fmt.Println("Hello World, the resource is here:",res)
}
