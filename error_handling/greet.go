package main

import (
	"fmt"
)

func greet(name string) error {
	if name == "" {
		return fmt.Errorf("Name cannot be empty")
	}
	return nil
}

func main() {
	name := "Lucky"
	err := greet(name)
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	fmt.Println("Hello there,", name)
}
