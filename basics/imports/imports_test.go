package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestImportsProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "Now you have 2.6457513110645907 problems.\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
