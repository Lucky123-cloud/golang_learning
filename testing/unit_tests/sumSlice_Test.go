package main

import "testing"

func TestSumSlice(t *testing.T) {
	// Test 1: normal slice
	res1 := SumSlice([]int{1, 2, 3})
	if res1 != 6 {
		t.Errorf("SumSlice([1, 2, 3]) = %d; want 6", res1)
	}

	// Test 2: empty slice
	res2 := SumSlice([]int{})
	if res2 != 0 {
		t.Errorf("SumSlice([]) = %d; want 0", res2)
	}

	// Test 3: negative numbers
	res3 := SumSlice([]int{5, -3, 2})
	if res3 != 4 {
		t.Errorf("SumSlice([5, -3, 2]) = %d; want 4", res3)
	}
}
