package main

import (
	"fmt"
)

func main() {
	//infinite loop does not have a condition, this is not advisable because it will always eat way our memory, we need to know that
	//for our exmaple we will put in a break keyword in order to avoid running entirely

	for {
		fmt.Println("Hello world, Lucky is the best Software engineer")
		break // if we dont add this, it will run
	}
}
