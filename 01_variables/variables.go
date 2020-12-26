package main

import "fmt"

var number, text = 9, "Lorem ipsum."

func printNumber() {
	const pi float64 = 3.141592653589793

	fmt.Println(number)
	fmt.Println(pi)
}

func printText(text string) {
	fmt.Println(text)
}

func main() {
	local := true

	printNumber()
	printText(text)
	fmt.Println(local)
}
