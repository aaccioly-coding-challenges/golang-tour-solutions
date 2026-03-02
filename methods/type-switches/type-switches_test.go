package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestTypeSwitchesProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "Twice 21 is 42\n\"hello\" is 5 bytes long\nI don't know about type bool!\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
