package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//solution 2 mutex
	count := int64(0)
	nGR, nIter := 10, 1_000

	var wg sync.WaitGroup

	wg.Add(nGR)
	for range nGR {
		go func() {
			defer wg.Done()
			for range nIter {
				//Solution2 : sync/atomic
				atomic.AddInt64(&count, 1)

				time.Sleep(time.Microsecond)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Count:", count)
}
