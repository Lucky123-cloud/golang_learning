package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	res1 := IsPrime(7)
	res2 := IsPrime(8)
	res3 := IsPrime(0)
	res4 := IsPrime(1)
	res5 := IsPrime(-1)

	if !res1 {
		t.Errorf("Error")
	}

	if res2 != false {
		t.Errorf("Error")
	}

	if res3 != false {
		t.Errorf("Error")
	}

	if res4 != false {
		t.Errorf("Error")
	}

	if res5 != false {
		t.Errorf("Error")
	}

}
