package main


import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("https://digibiashara.com")
	defer resp.Body.Close()

	file, _ := os.Create("page.html")
	defer file.Close()

	io.Copy(file, resp.Body)
}

