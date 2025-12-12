//count how many times a word appears in a sentence - we need to first of all to break the text/line into tokens 


//here is the implementation


package main


import (
	"fmt"
	"strings"
)


//WordCount counts the number of times a word appears in a full sentence
//it returns a map of words as a string and the number of times they appear as integers
func WordCount(text string) map[string]int{
	countWord := make(map[string]int)
	words := strings.Fields(text)

	for _, w := range words {
		countWord[w]++
	}
	return countWord
}

func main() {
	lines := "Hello World I am Lucky Baraka, and I am Lucky to be a go programmer because go is one of the best in the word"
	res := WordCount(lines)
	fmt.Println(res)
}

