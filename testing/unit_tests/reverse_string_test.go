package main

import (
	"testing"
)

// func TestReverse_string(t *testing.T) {
// 	res := Reverse_string("hello")
// 	res2 := Reverse_string("")
// 	res3 := Reverse_string("level")

// 	if res != "olleh" {
// 		t.Errorf("There is an error, expected %s but got %s", "olleh", res)
// 	}

// 	if res2 != "" {
// 		t.Errorf("There is an error, expected %s but got %s", "", res2)
// 	}

// 	if res3 != "level" {
// 		t.Errorf("There is an error, expected %s but got %s", "level", res3)
// 	}

// }

// we are now testing using table driven test

func TestReverse_String(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"'hello' becomes 'olleh'", "hello", "olleh"},
		{"'' becomes ''", "", ""},
		{"'level' becomes 'level'", "level", "level"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Logf("strating with %s expecting %s", tt.input, tt.want)
			got := Reverse_string(tt.input)
			if got != tt.want {
				t.Errorf("Testing case failed, expected %s, got %s", tt.want, got)
			}
		})
	}
}
