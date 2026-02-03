package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestNumericConstantsProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "21\n0.2\n1.2676506002282295e+29\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
