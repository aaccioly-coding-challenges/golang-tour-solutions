package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now()
	switch time.Saturday {
	case today.Weekday() + 0:
		fmt.Println("Today.")
	case today.Weekday() + 1:
		fmt.Println("Tomorrow.")
	case today.Weekday() + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
