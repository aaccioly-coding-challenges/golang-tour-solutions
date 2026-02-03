package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestArrayProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "Hello World\n[Hello World]\n[2 3 5 7 11 13]\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
