package main

import (
	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
	"testing"
)

func TestForIsGoWhileProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "1024\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
