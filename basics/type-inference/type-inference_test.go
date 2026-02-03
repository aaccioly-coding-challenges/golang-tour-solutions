package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestTypeInferenceProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "v is of type int\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
