package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestRangeContinuedProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "1\n2\n4\n8\n16\n32\n64\n128\n256\n512\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}

}
