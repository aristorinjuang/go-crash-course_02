package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		"John Doe",
		1e1,
	}
	p2 := person{
		age: 999,
	}

	p1.age += 8

	fmt.Println(p1)
	fmt.Println(p1.name, "is", p1.age, "years old.")

	fmt.Println(p2)
}
