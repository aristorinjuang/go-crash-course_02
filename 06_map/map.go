package main

import "fmt"

type location struct {
	latitude, longitude float64
}

func main() {
	offices := map[string]location{
		"Apple": {
			37.33199, -122.03089,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	fmt.Println(offices["Apple"])
	fmt.Println("Google is at:", offices["Google"].latitude, offices["Google"].longitude)
}
