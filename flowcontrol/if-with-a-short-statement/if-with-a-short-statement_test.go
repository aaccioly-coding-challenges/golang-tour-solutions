package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestIfWithAShortStatementProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "9 20\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
