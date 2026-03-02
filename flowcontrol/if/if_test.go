package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestIfProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "1.4142135623730951 2i\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
