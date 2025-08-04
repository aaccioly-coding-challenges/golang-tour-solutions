package main

import (
	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
	"testing"
)

func TestIfProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "1.4142135623730951 2i\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
