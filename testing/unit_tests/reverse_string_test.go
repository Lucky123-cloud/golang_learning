package main

import "testing"

func TestReverse_string(t *testing.T) {
	res := Reverse_string("hello")
	res2 := Reverse_string("")
	res3 := Reverse_string("level")

	if res != "olleh" {
		t.Errorf("There is an error, expected %s but got %s", "olleh", res)
	}

	if res2 != "" {
		t.Errorf("There is an error, expected %s but got %s", "", res2)
	}

	if res3 != "level" {
		t.Errorf("There is an error, expected %s but got %s", "level", res3)
	}

}
