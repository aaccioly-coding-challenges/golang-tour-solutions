package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestVariablesWithInitializersProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "1 2 true false no!\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
