package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestShortVariableDeclarationProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "1 2 3 true false no!\n"
	if output != expected {
		t.Errorf("Expected output  %q, got %q", expected, output)
	}
}
