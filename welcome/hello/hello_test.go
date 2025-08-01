package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestHelloProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "Hello, 世界\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
