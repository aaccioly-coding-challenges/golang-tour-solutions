package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestVariablesProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "0 false false false\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
