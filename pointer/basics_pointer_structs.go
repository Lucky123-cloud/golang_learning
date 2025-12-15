package main


import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func UpdateAge(u *User) {
	u.Age = 30
}

func main() {
	person := User{Name: "Lucky", Age: 20}
	person2 := User{Name: "Alex", Age: 25}
	p := &person2
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	UpdateAge(&person)
	fmt.Println(person.Age) //still will be 30
}

