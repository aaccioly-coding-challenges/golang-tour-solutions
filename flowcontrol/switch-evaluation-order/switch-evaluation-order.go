package main

import (
	"fmt"
	"time"
)

func whenIsSaturday(t time.Time) string {
	switch time.Saturday {
	case t.Weekday() + 0:
		return "Today."
	case t.Weekday() + 1:
		return "Tomorrow."
	case t.Weekday() + 2:
		return "In two days."
	default:
		return "Too far away."
	}
}

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now()
	fmt.Println(whenIsSaturday(today))
}
