package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestMethodsWithPointerReceiversProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "Before scaling: &{X:3 Y:4}, Abs: 5\nAfter scaling: &{X:15 Y:20}, Abs: 25\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
