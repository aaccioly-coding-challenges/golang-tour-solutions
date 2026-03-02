package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestConstantsProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "Hello 世界\nHappy 3.14 Day\nGo rules? true\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
