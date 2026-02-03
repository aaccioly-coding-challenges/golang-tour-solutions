package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestFunctionValuesProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "13\n5\n81\n"

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
