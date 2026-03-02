package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestRot13Program(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "You cracked the code!"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
