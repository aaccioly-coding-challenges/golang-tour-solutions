package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestExportedNamesProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "3.141592653589793\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
