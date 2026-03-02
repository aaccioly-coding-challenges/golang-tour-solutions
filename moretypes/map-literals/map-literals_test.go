package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestMapLiteralsProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
