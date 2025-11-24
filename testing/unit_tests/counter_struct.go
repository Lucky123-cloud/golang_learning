package main

import (
	"fmt"
)

type Counter struct {
	Value int
}

func (c *Counter) Increament(n int) {
	c.Value += n
}

func main() {
	counting := Counter{Value: 20}
	fmt.Println(counting)

	counting.Increament(100)
	fmt.Println(counting)

}
