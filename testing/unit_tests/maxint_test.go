package main

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	// test scenario one
	res := MaxInt(3, 7)

	if res != 7 {
		t.Errorf("Wrong result, expected %d, got %d", 7, res)
	}

	//scenario two
	res2 := MaxInt(10, -5)
	if res2 != 10 {
		t.Errorf("Wrong result, expected %d, got %d", 10, res2)
	}

	// scenario 3
	res3 := MaxInt(4, 4)
	if res3 != 4 {
		t.Errorf("Wrong anser, expected, %d, got %d", 4, res3)
	}

}
