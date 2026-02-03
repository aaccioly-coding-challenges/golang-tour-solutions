package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestExportedNamesProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "3.141592653589793\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
