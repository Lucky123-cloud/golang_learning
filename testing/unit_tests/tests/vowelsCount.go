package main

import (
	"strings"
)

func CountVowels(s string) int {
	s = strings.ToLower(s)
	count := 0

	for _, n := range s {
		if n == 'a' || n == 'e' || n == 'i' || n == 'o' || n == 'u' {
			count++
		}
	}
	return count

}
