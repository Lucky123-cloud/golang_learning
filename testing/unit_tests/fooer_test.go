package main

import (
	"testing"
)

func TestFooer(t *testing.T) {
	result := Fooer(3)

	if result != "Foo" {
		t.Errorf("The Test Failed, expected %s but got %s", "Foo", result)
	}

}
