package main

import (
	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
	"testing"
)

func TestImportsProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "Now you have 2.6457513110645907 problems.\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
