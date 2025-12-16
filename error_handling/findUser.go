package main


import (
	"fmt"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

func findUser(id int) error {
	if id == 0 {
		return ErrUserNotFound
	}
	return nil
}

func main() {
	var id int
	err := findUser(id)
	if err == ErrUserNotFound{
		fmt.Println("Error fetching user: ", err)
	}
}
