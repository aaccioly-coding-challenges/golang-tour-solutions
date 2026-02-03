// Package testutils provides common testing utilities for capturing program output
package testutils

import (
	"bytes"
	"os"
)

// CaptureOutput captures stdout while running the provided function and returns the output
func CaptureOutput(f func()) string {
	output, _ := CaptureOutputWithPanic(f)
	return output
}

// CaptureOutputWithPanic captures stdout while running the provided function.
// It returns the captured output and any panic that occurred.
func CaptureOutputWithPanic(f func()) (output string, recovered any) {
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

	f()
	return
}
