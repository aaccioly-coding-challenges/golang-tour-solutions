package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestStructPointesProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "{1000000000 2}\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
