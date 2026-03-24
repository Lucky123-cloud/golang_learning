package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// func main() {
// 	var s, sep string
// 	for i := 0; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	for i, arg := range os.Args {
// 		fmt.Printf("index: %d,  Arguments: %s\n", i, arg)
// 	}
// }

// measuring time and perfomance
func main() {
	//inefficient code
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	elapsed := time.Since(start)
	fmt.Println("inefficient time: ", elapsed)

	//effience version
	start = time.Now()
	result := strings.Join(os.Args, " ")
	fmt.Println(result)

	elapsed = time.Since(start)
	fmt.Println("efficient time: ", elapsed)
}
