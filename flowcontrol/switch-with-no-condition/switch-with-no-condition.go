package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	greet(t)
}

func greet(t time.Time) {
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
