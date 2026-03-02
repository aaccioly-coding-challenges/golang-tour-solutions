package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestImportsProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "Now you have 2.6457513110645907 problems.\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
