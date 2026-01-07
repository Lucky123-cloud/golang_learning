package main

import (
	"testing"
	"fmt"
)

func TestRomanToInt(t *testing.T) {
	test := []struct {
		name string
		input string
		expected int
	}{
		{
			"Simple number",
			"IX",
			9,
		},
		{
			"mixed number",
			"LVIII",
			58,
		},
		{
			"complex numbers",
			"MCMXCIV",
			1994,
		},
		{
			"no input",
			"",
			0,
		},
		{
			"subtractive notation",
			"IV",
			4,
		},
		{
			"40 and 90",
			"XLIX",
			49,
		},
		{
			"400 and 900",
			"CMXC",
			990,
		},
	}

	for _, tt := range test{
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.input)
			if got != tt.expected{
				fmt.Errorf("Wanted %d, got %d", tt.expected, got)
			}
		})
	}
}
