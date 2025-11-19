package main

import (
	"fmt"
	"math"
)

const n = "Lucky is the best Engineer"

func main() {

	fmt.Println(n)

	const num = 55555477464722

	// make d a runtime float64 variable
	d := 3e20 / float64(num)

	fmt.Println(d)

	fmt.Println(int64(d)) // now works

	fmt.Println(math.Sin(float64(num)))
}
