package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Server")
	})
	fmt.Println("Server running in port 5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Error occured while starting a server:", err)
	}
}
