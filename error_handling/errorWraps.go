package main


import (
	"fmt"
	"errors"
)

var ErrUnauthorized = errors.New("Unauthorized")

func authenticate(ok bool) error {
	if !ok {
		return ErrUnauthorized
	}
	return nil
}

func login(ok bool) error {
	err := authenticate(ok)
	if err != nil {
		return fmt.Errorf("Login failed: %w", err)
	}
	return nil
}

func main() {
	err := login(false)

	if errors.Is(err, ErrUnauthorized) {
		fmt.Println("Access denied: ", err)
		return
	}
	fmt.Println("Login sucessful")
}

