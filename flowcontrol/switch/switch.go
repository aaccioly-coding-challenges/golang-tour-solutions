package main

import (
	"fmt"
	"runtime"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Print("Go runs on ")
	os := runtime.GOOS
	fmt.Println(switchOs(os))
}

func switchOs(os string) string {
	switch os {
	case "darwin":
		return "macOS."
	case "linux":
		return "Linux."
	default:
		// Freebsd, Openbsd,
		// Plan9, Windows...
		return fmt.Sprintf("%s.", cases.Title(language.Und).String(os))
	}
}
