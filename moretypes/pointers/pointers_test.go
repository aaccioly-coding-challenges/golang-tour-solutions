package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestPointersProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "42\n21\n73\n"

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
