package main

import (
	"strconv"
)

func Fooer(num int) string {
	isFoo := (num % 3) == 0

	if isFoo {
		return "Foo"
	}

	return strconv.Itoa(num)
}
