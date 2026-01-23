// Package testutils provides common testing utilities for capturing program output
package testutils

import (
	"bytes"
	"os"
)

// CaptureMainOutput captures stdout while running the provided main function and returns the output
func CaptureMainOutput(mainFunc func()) string {
	output, _ := CaptureMainOutputWithPanic(mainFunc)
	return output
}

// CaptureMainOutputWithPanic captures stdout while running the provided main function.
// It returns the captured output and any panic that occurred.
func CaptureMainOutputWithPanic(mainFunc func()) (output string, recovered any) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run main function and recover from panic
	defer func() {
		// Restore stdout
		_ = w.Close()
		os.Stdout = old

		// Read captured output
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)
		output = buf.String()
		recovered = recover()
	}()

	mainFunc()
	return
}
