package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestTypeAssertionsProgram(t *testing.T) {
	output, recovered := testutils.CaptureMainOutputWithPanic(main)
	expectedOutput := "hello\nhello true\n0 false\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
	if recovered == nil {
		t.Errorf("Expected panic, got none")
	}
}
