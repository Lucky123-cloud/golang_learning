package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// initial guess
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: z = %.6f\n", i+1, z)
	}
	return z

}

func main() {
	numbers := []float64{1, 2, 3, 4, 9, 16, 20}

	for _, x := range numbers {
		fmt.Printf("\nComputing sqrt(%.2f)\n", x)
		z := Sqrt(x)
		fmt.Printf("My Sqrt: %.6f | math.Sqrt: %.6f\n", z, math.Sqrt(x))
	}
}
