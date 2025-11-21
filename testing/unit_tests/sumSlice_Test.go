package main

import (
	"testing"
)

func TestSumSlice(t *testing.T) {
	res := SumSlice([]int{1, 2, 3})

	if res != 6 {
		t.Errorf("Wrong Answer")
	}

}
