package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	goBirthday := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("If this was the real playground is would be always", goBirthday)
}
