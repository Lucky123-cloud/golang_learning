package main

import (
	"testing"
)

// func TestCountVowels(t *testing.T) {
// 	// Test 1: normal sentence
// 	res1 := CountVowels("Hello World")
// 	if res1 != 3 {
// 		t.Errorf("CountVowels(\"Hello World\") = %d; want 3", res1)
// 	}

// 	// Test 2: all vowels
// 	res2 := CountVowels("aeiou")
// 	if res2 != 5 {
// 		t.Errorf("CountVowels(\"aeiou\") = %d; want 5", res2)
// 	}

// 	// Test 3: empty string
// 	res3 := CountVowels("")
// 	if res3 != 0 {
// 		t.Errorf("CountVowels(\"\") = %d; want 0", res3)
// 	}

// 	// Test 4: mixed case
// 	res4 := CountVowels("Golang")
// 	if res4 != 2 {
// 		t.Errorf("CountVowels(\"Golang\") = %d; want 2", res4)
// 	}

// 	// Test 5: no vowels
// 	res5 := CountVowels("bcdfg")
// 	if res5 != 0 {
// 		t.Errorf("CountVowels(\"bcdfg\") = %d; want 0", res5)
// 	}
// }

// let's now test using TDT

func TestCountVowelsTableDrivenTest(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"In 'Hello World', we have 3 vowels", "Hello World", 3},
		{"'aeiou' has 5 vowels", "aeiou", 5},
		{" '' gives 0 vowels", "", 0},
		{" 'Golang' has 2 vowels", "Golang", 2},
		{"'bdgfhgd' has no vowel count hence 0", "bdfhg", 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Logf("Starting the test with %s, expecting %d", tt.input, tt.want)
			got := CountVowels(tt.input)
			if got != tt.want {
				t.Errorf("Test failed!, expected %d, got %d", tt.want, got)
			}
		})
	}
}
