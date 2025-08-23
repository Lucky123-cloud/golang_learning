package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for j := 0; j <= 100; j += 10 {
		fmt.Println(j)
	}

	for k := 0; k < 5; k++ {
		if k == 3 {
			continue
		} else if k == 4 {
			break
		}
		fmt.Println(k)

	}

	adj := [3]string{"red", "big", "tasty"}
	fruits := [3]string{"apple", "banana", "melon"}
	for l := 0; l < len(adj); l++ {
		for m := 0; m < len(fruits); m++ {
			fmt.Println(adj[l], fruits[m])
		}
	}

	//using range in for loop
	numbers := [5]int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("%v\t%v\n", index, value)
	}

	//using _ to ignore index or value
	for _, value := range numbers {
		fmt.Printf("%v\n", value)
	}

	//using _ to ignore index or value
	for index, _ := range numbers {
		fmt.Printf("%v\n", index)
	}

}
