package main

import "fmt"

func main() {
	number := 9
	pointer := &number

	fmt.Println(*pointer)
	fmt.Println(pointer)

	*pointer = number + 1

	fmt.Println(number)
}
