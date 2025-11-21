package main

import (
	"testing"
)

// func TestIsPrime(t *testing.T) {
// 	res1 := IsPrime(7)
// 	res2 := IsPrime(8)
// 	res3 := IsPrime(0)
// 	res4 := IsPrime(1)
// 	res5 := IsPrime(-1)

// 	if !res1 {
// 		t.Errorf("Error")
// 	}

// 	if res2 != false {
// 		t.Errorf("Error")
// 	}

// 	if res3 != false {
// 		t.Errorf("Error")
// 	}

// 	if res4 != false {
// 		t.Errorf("Error")
// 	}

// 	if res5 != false {
// 		t.Errorf("Error")
// 	}

// }

// Let's use the table driven test

func TestIsPrimeTableDrivenTest(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"7 is Prime number", 7, true},
		{"8 is not prime", 8, false},
		{"0 is not prime", 0, false},
		{"1 is not prime", 1, false},
		{"-1 is not prime", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPrime(tt.input)
			if got != tt.want {
				t.Errorf("Expected %t, got %t", tt.want, got)
			}
		})
	}
}
