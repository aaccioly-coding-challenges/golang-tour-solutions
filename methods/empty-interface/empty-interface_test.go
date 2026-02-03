package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestEmptyInterfaceProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "(<nil>, <nil>)\n(42, int)\n(hello, string)\n"

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
