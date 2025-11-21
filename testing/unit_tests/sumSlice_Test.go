package main

import (
	"testing"
)

// func TestSumSlice(t *testing.T) {
// 	// Test 1: normal slice
// 	res1 := SumSlice([]int{1, 2, 3})
// 	if res1 != 6 {
// 		t.Errorf("SumSlice([1, 2, 3]) = %d; want 6", res1)
// 	}

// 	// Test 2: empty slice
// 	res2 := SumSlice([]int{})
// 	if res2 != 0 {
// 		t.Errorf("SumSlice([]) = %d; want 0", res2)
// 	}

// 	// Test 3: negative numbers
// 	res3 := SumSlice([]int{5, -3, 2})
// 	if res3 != 4 {
// 		t.Errorf("SumSlice([5, -3, 2]) = %d; want 4", res3)
// 	}
// }

// let's test using Table Driven Test

func TestSumSliceTableDrivenTest(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"sum of []int{1, 2, 3} is 6", []int{1, 2, 3}, 6},
		{"sum of []int{} is 0", []int{}, 0},
		{"sum of []int{1, -2, 5} is 4", []int{1, -2, 5}, 4},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Logf("Starting with %d expecting %d", tt.input, tt.want)
			got := SumSlice(tt.input)
			if got != tt.want {
				t.Errorf("Testing case failed, expected %d, got %d", tt.want, got)
			}
		})
	}
}
