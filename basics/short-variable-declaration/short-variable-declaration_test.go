package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestShortVariableDeclarationProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "1 2 3 true false no!\n"
	if output != expected {
		t.Errorf("Expected output  %q, got %q", expected, output)
	}
}
