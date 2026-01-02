package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	count := 0

	nGR, nIter := 10, 1_000
	var wg sync.WaitGroup

	wg.Add(nGR)
	for range nGR {
		go func() {
			for range nIter {
				mu.Lock()
				count++
				mu.Unlock()
				/*
					 * Fetch count
						* increament count
						* store count
				*/
				time.Sleep(time.Microsecond)
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Count:", count)
}

/*
 * go run -race count.go
 * "-race" is supported by
 * -run
 * -build
 * -test
 *
 * Rule of thumb: use "go test -race"
 */
