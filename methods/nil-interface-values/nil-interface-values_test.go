package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestNilInterfaceValuesProgram(t *testing.T) {
	output, recovered := testutils.CaptureMainOutputWithPanic(main)
	expectedOutput := "(<nil>, <nil>)\n"
	if output != expectedOutput {
		t.Errorf("Expected output to be '%s', but got '%s'", expectedOutput, output)
	}
	if recovered == nil {
		t.Error("Expected program to panic, but it didn't")
	}
}
