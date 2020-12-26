package main

import "fmt"

func typical() {
	for i := 1; i <= 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		fmt.Print("\n")
	}
}

func while() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func main() {
	typical()
	while()
}
