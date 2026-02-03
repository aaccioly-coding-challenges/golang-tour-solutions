package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestFibonacciClosureProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "1\n1\n2\n3\n5\n8\n13\n21\n34\n55\n"

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
