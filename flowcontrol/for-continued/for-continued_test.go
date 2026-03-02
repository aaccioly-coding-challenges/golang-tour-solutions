package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestForContinuedProgram(t *testing.T) {
	output := capture.CaptureOutput(main)

	// The program outputs the sum of numbers until it reaches or exceeds 1000
	expectedOutput := "1024\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}
