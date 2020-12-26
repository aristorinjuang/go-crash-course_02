package main

import "fmt"

func main() {
	hello := [2]string{
		"Hello",
		"World!",
	}
	primes := []int{2, 3, 5}

	fmt.Println(hello)
	fmt.Println(primes)

	fmt.Println(hello[0])
	fmt.Println(primes[1:])
	fmt.Println(primes[1:2])
}
