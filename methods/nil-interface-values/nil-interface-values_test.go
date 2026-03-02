package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestNilInterfaceValuesProgram(t *testing.T) {
	output, recovered := capture.CaptureOutputWithPanic(main)
	expectedOutput := "(<nil>, <nil>)\n"
	if output != expectedOutput {
		t.Errorf("Expected output to be '%s', but got '%s'", expectedOutput, output)
	}
	if recovered == nil {
		t.Error("Expected program to panic, but it didn't")
	}
}
