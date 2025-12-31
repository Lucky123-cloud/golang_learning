package main

import (
	"fmt"
	"runtime"
)

type Person struct {
	Name string
	Age  int
}

func UpdatePerson(p *Person) {
	p.Name = "Updated Name"
	p.Age = 30
}

func allocateLargeArray() {
	big := make([]byte, 100_000_000)
	for i := range big {
		big[i] = 1
	}
	fmt.Println("Large array allocated")
}

func printMemStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("\n[%s]\n", label)
	fmt.Printf("Alloc = %v MB\n", m.Alloc/1024/1024)
	fmt.Printf("Sys = %v MB\n", m.Sys/1024/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}

func main() {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "charlie", Age: 40},
	}

	fmt.Println("Before update:", people)
	UpdatePerson(&people[1])
	fmt.Println("After update:", people)

	printMemStats("Before Allocation")
	allocateLargeArray()
	printMemStats("After allocation")

	runtime.GC()
	printMemStats("After forced GC")
}
