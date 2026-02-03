package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestIfWithAShortStatementProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "9 20\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
