package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestTypeConversionsProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "3 4 5\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
