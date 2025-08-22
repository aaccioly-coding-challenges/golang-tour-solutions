package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestMutatingMapsProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "The value: 42\nThe value: 48\nThe value: 0\nThe value: 0 Present? false\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
