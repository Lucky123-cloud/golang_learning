package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func multiply(a int, b int) (result int) {
	result = a * b
	return result
}

func myFunc(x int, y string) (result int, text1 string) {
	result = x + x
	text1 = y + " World!"
	return
}

func divide(a int, b int, name string) (result int, say string) {
	result = a / b
	say = "Hello " + name
	return
}

func main() {
	res := add(10, 20)
	fmt.Println("Result is:", res)
	fmt.Println("Multiply is:", multiply(5, 6))
	a, b := myFunc(10, "Hello")
	fmt.Println(myFunc(10, "Hello"))
	fmt.Println(a, b)
	_, d := divide(20, 5, "Shangwe")
	c, _ := divide(20, 5, "Shangwe")
	fmt.Println(d)
	fmt.Println(c)

}
