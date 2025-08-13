package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestStructLiteralsProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "{1 2} &{1 2} {1 0} {0 0}\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
