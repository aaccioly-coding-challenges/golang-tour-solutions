package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestRangeProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "2**0 = 1\n2**1 = 2\n2**2 = 4\n2**3 = 8\n2**4 = 16\n2**5 = 32\n2**6 = 64\n2**7 = 128\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
