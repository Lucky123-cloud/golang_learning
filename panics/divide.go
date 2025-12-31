// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(safeDivide(6, 3))
// 	fmt.Println(safeDivide(6, 0))
// }

// func safeDivide(a, b int) (q int, err error) {
// 	defer func() {
// 		e := recover()
// 		if e != nil {
// 			err = fmt.Errorf("Error: %v", e)
// 		}
// 	}()

// 	return divide(a, b), nil
// }

// func divide(a, b int) int {
// 	if b == 0 {
// 		panic("divisor cannot be zero(0)")
// 	}
// 	return a / b
// }
