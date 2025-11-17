package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestIndirectionValuesProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "5\n5\n5\n5\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}

}
