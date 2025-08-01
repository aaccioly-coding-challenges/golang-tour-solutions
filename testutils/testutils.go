// Package testutils provides common testing utilities for capturing program output
package testutils

import (
	"bytes"
	"os"
)

// CaptureMainOutput captures stdout while running the provided main function and returns the output
func CaptureMainOutput(mainFunc func()) string {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run main function
	mainFunc()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}