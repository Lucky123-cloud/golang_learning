package main


import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", w http.ResponseWriter, r *http.RequestWriter) {
		return fmt.Sprintf(w, "Hello World from Kenya") 
	}
	fmt.Println("Server Running from port 5000")
	http.ListenAndServe(":5000")
}
