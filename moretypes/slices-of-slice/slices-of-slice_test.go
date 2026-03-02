package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestSlicesOfSlices(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "X _ X\nO _ X\n_ _ O\n"

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
