package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestMultipleResultsProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "world hello\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
