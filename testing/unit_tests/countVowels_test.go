package main

import (
	"testing"
)

func TestCountVowels(t *testing.T) {
	// Test 1: normal sentence
	res1 := CountVowels("Hello World")
	if res1 != 3 {
		t.Errorf("CountVowels(\"Hello World\") = %d; want 3", res1)
	}

	// Test 2: all vowels
	res2 := CountVowels("aeiou")
	if res2 != 5 {
		t.Errorf("CountVowels(\"aeiou\") = %d; want 5", res2)
	}

	// Test 3: empty string
	res3 := CountVowels("")
	if res3 != 0 {
		t.Errorf("CountVowels(\"\") = %d; want 0", res3)
	}

	// Test 4: mixed case
	res4 := CountVowels("Golang")
	if res4 != 2 {
		t.Errorf("CountVowels(\"Golang\") = %d; want 2", res4)
	}

	// Test 5: no vowels
	res5 := CountVowels("bcdfg")
	if res5 != 0 {
		t.Errorf("CountVowels(\"bcdfg\") = %d; want 0", res5)
	}
}
