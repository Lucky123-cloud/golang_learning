package main

import (
	"testing"
)

// func TestFooer(t *testing.T) {
// 	result := Fooer(3)

// 	if result != "Foo" {
// 		t.Errorf("The Test Failed, expected %s but got %s", "Foo", result)
// 	}

// }

// we want to now use table driven test
func TestFooerTableDrivenTest(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should be Foo", 0, "Foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Starting the test with %d, expecting %s", tt.input, tt.want)
			got := Fooer(tt.input)
			if got != tt.want {
				t.Errorf("Got %s, expected %s", got, tt.want)
			}
		})
	}
}
