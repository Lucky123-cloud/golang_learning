package main
import (
	"fmt"
	"time"
	"net/http"
)

func main() {
	urls := []string{
		"https://go.dev",
		"https://ardanlabs.com",
		"https://ibm.com/no/such/page",
	}

	start := time.Now()
	/*
	for _, url := range urls {
		stat, err := urlCheck(url)
		fmt.Printf("%q: %d (%v)\n", url, stat, err)
	}
	*/
	fanOutResult(urls)
	duration := time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
}

func fanOutResult(urls [] string) {
	type result struct {
		url string
		status int
		err error
	}
	ch := make(chan result)

	//fan-out
	for _, url := range urls {
		go func() {
			r := result{url: url}
			defer func() {ch <- r}()

			r.status, r.err = urlCheck(url)
		}()
	}

	//collect results
	for range urls {
		r := <- ch
		fmt.Printf("%q: %d (%v)\n", r.url, r.status, r.err)
	}
}

func urlCheck(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
