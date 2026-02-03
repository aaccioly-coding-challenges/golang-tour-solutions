package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestIfAndElseProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "27 >= 20\n9 20\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}

}
