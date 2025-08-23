package main

import (
	"fmt"
)

func familyName(fname string, age int) {
	fmt.Println("Hello", fname, "Baraka", "age is", age)
}

func main() {
	familyName("Shangwe", 10)
	familyName("John", 12)
	familyName("Mark", 15)
}
