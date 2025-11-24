package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Increament() {
	c.Value++
}

type Increamental interface {
	Increament()
}

func main() {
	c := &Counter{Value: 20}
	var inc Increamental = c

	inc.Increament()
	fmt.Println(c.Value)
}
