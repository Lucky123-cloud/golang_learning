package main

import (
	"fmt"
)

func main() {
	// Define a handler function
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, World! This is a basic Go server 🚀")
	// })

	// // Start the server on port 8080
	// fmt.Println("Server is running on http://localhost:8080")
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }
	//
	i, j := 32, 34
	fmt.Println(&i, &j)

	p := &i
	fmt.Println(*p)
	*p = 30
	fmt.Println(*p)

	*p = *p / 5
	fmt.Println(i)
}
