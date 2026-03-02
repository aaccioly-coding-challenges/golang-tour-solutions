package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestStringerProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
