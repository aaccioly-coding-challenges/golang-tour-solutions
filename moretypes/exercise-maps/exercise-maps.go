package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	frequency := make(map[string]int)
	for _, word := range strings.Fields(s) {
		frequency[word]++
	}
	return frequency
}

func main() {
	wc.Test(WordCount)
}
