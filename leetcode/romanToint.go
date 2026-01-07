package main


func romanToInt(s string) int {
	//Map Roman numerals to integers
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	//Iterate using runes to avoid byte/run issues
	for i := 0; i < n; i++ {
		current := values[rune(s[i])]
		if i < n-1 && current < values[rune(s[i+1])] {
			total -= current
		}else {
			total += current
		}
	}
	return total
}


