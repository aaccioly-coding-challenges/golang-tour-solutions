package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestStructFieldsProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "4\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
