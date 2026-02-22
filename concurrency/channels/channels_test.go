package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestChannelsProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "-5 17 12\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
