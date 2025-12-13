package main


import (
	"fmt"
)
func CircleStats(radius float64) (area, circumference float64) {
	const PI = 3.142
	area = PI * radius * radius
	circumference = PI * radius * 2
	return
}

func main() {
	a, c := CircleStats(23.5)
	fmt.Println("Area of the circle is:", a)
	fmt.Println("Circumference is:", c)
}
