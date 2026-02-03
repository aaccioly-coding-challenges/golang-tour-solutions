package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSliceLengthAndCapacityProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "len=6 cap=6 [2 3 5 7 11 13]\nlen=0 cap=6 []\nlen=4 cap=6 [2 3 5 7]\nlen=2 cap=4 [5 7]\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
