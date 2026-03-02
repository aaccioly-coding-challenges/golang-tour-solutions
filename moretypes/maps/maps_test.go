package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestMapsProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "{40.68433 -74.39967}\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}

}
