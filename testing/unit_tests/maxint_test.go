package main

import (
	"testing"
)

// func TestMaxInt(t *testing.T) {
// 	// test scenario one
// 	res := MaxInt(3, 7)

// 	if res != 7 {
// 		t.Errorf("Wrong result, expected %d, got %d", 7, res)
// 	}

// 	//scenario two
// 	res2 := MaxInt(10, -5)
// 	if res2 != 10 {
// 		t.Errorf("Wrong result, expected %d, got %d", 10, res2)
// 	}

// 	// scenario 3
// 	res3 := MaxInt(4, 4)
// 	if res3 != 4 {
// 		t.Errorf("Wrong anser, expected, %d, got %d", 4, res3)
// 	}

// }

// we want to test using Table Driven Test

func TestMaxIntTableDrivenDevelopment(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"7 is greater than 3", 3, 7, 7},
		{"10 is greater than -5", 10, -5, 10},
		{"4 is equal to 4", 4, 4, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("starting with %d  and %d expecting %d", tt.a, tt.b, tt.want)
			got := MaxInt(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Test did not pass, expected %d, got %d", tt.want, got)
			}
		})
	}
}
