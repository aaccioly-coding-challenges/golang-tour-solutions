package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSelectProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "0\n1\n1\n2\n3\n5\n8\n13\n21\n34\nquit\n"

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
