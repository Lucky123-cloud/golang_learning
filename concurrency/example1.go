package main

import (
	"fmt"
	"time"
)

func main() {
	//simple goroutine
	go fmt.Println("goroutine")
	fmt.Println("Main")

	//Goroutine in loop (correct variable capture)
	for i := 0; i < 3; i++ {
		i := i //capture the loop variable
		go func () {
			fmt.Println("goroutine:", i)
		}()
	}

	time.Sleep(10 * time.Millisecond)

	//channel example
	ch := make(chan int)
	go func() {
		ch <- 7 //sending data over channel to goroutine
	}()

	v := <- ch //receiving data
	fmt.Println(v)

	//Sleep sort example 
	fmt.Println(sleepSort([]int{20, 30, 10})) //[10 20 30]

	go func(){
		for i := range 4 {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(">>", v)
	}

	v, ok := <-ch
	fmt.Println("Closed:", v, "ok", ok)
	/*
	The "for range" above does
	for {
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(">>", v)
	}
	*/
}

/*
Algorithm (sleepSort)
- For every value n in values:
	- start a goroutine
	- sleep for n milliseconds
	- send n over a channel
- Collect all values from the channel
- Return them in the order received
- Receive from a closed channel will return zero value with no block
*/


func sleepSort(values []int) []int {
	ch := make(chan int)

	for _, n := range values {
		n := n //capture value
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}()
	}

	var out []int
	for range values {
		n := <-ch
		out = append(out, n)
	}
	return out
}

/*
Channel Semantics

- sending and receving on an unbuffered channel blocks
- Sends wait for a receiver
- receive waits for a sender
- Delivery is guarnteed
- Receive from a closed channel will return zero value with no block
	- use "comma ok" to check if channel was closed
- Send to a closed channel will panic
- Closing a closed or nil channel will panic
- send/receive to a nil channel will block forever
*/

