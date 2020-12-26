package main

import (
	"fmt"
	"time"
)

func noCondition() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func typical(day string) {
	switch day {
	case "Monday":
		fmt.Println("Mathematics")
	case "Wednesday":
		fmt.Println("Computer Science")
	case "Friday":
		fmt.Println("Physics")
	default:
		fmt.Println("Off")
	}
}

func main() {
	typical("Today")
	noCondition()
}
