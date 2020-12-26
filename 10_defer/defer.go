package main

import "fmt"

func main() {
	defer fmt.Println("World!")

	fmt.Println("Hello...")

	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
}
