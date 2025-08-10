package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestDeferProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "hello\nworld\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
