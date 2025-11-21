package main

func Reverse_string(str string) string {
	var newStr string

	for _, char := range str {
		newStr = string(char) + newStr
	}

	return newStr

}
