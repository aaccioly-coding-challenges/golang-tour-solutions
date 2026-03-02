package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestImagesProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "(0,0)-(100,100)\n0 0 0 0\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
