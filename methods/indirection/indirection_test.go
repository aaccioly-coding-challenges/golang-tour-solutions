package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestMethodsAndPointerIndirectionProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "{60 80} &{96 72}\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
