package main

import (
	"strings"
)

func CountVowels(s string) int {
	count := 0

	s = strings.ToLower(s)

	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			count++
		}
	}
	return count
}
