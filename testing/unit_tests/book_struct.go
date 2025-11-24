package main

import (
	"fmt"
)

type Book struct {
	Title string
	Pages int
}

// creating a method
func (b Book) Summary() string {
	return fmt.Sprintf("Title: %s, Pages: %d", b.Title, b.Pages)
}

func (b *Book) UpdatePages(p int) int {
	b.Pages = p
	return p
}

func main() {
	myBook := Book{Title: "Pragmatic Programmer", Pages: 50}

	fmt.Println(myBook)
	fmt.Println(myBook.Summary())
	myBook.UpdatePages(120)
	fmt.Println(myBook.Summary())

}
