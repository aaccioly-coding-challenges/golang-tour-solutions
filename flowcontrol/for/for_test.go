package main

import (
	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
	"testing"
)

func TestForProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)

	// The program outputs the sum of numbers from 0 to 9, which is 45
	expectedOutput := "45\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}
