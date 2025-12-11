//count how many times a word appears

package main


import "fmt"
import "strings"

func CountWords(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(text)

	for _, w := range words {
		wordCount[w]++
	}
	return wordCount

}

func main() {
	text := "go is fun and go is powerful"
	counts := CountWords(text)
	fmt.Println(counts)
}


