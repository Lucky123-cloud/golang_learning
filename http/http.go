package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World from Lucky Baraka, the best Golang Software Engineer in Savannah Informatics")
	})

	fmt.Println("Hello World, we are running on server http:localhost:5000")
	http.ListenAndServe(":5000", nil)
}
