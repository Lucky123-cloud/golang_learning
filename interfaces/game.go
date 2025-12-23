package main

import "fmt"

func main() {
	var i Item
	fmt.Printf("i: %#v\n", i)

	i = Item{10, 20} //must specifcy all the fields
	fmt.Printf("i: %#v\n", i)

	i = Item{
		Y: 22,
		// X: 11,
	}
	fmt.Printf("i: %#v\n", i)

	fmt.Println(NewItem(10, 2000))
	fmt.Println(NewItem(10, 20))

	/* We use %#v for debugging/logging
	a, b := 1, "1"
	fmt.Printf("a=%v, b=%v\n", a, b)
	fmt.Printf("a=%#v, b=%#v\n", a, b)
	*/
}

/*
Types of new or factory function

func NewItem(x, y int) Item
func NewItem(x, y int) *Item
func NewItem(x, y int) (Item, error)
func NewItem(x, y int) (*Item, error)
*/

func NewItem(x, y int) (Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return Item{}, fmt.Errorf("%d/%d out of bounds %d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}

	return i, nil
}

const (
	maxX = 600
	maxY = 400
)

type Item struct {
	X int
	Y int
}

/*when to use pointer semantics
- When you have a lock field
- If you need to mutate the struct
- Decoding/unmarshalling
*/

/*
- Dont use any :) in interfaces (empty interfaces)
- Exceptions:
	- Serialization
	- Printing
*/
