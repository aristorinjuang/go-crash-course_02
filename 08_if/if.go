package main

import (
	"fmt"
)

func isEven(number int) {
	if number%2 == 0 {
		fmt.Println(number, "is even.")
	} else {
		fmt.Println(number, "is odd.")
	}
}

func main() {
	isEven(1)
}
