package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestStackingDefersPrograms(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "counting\ndone\n9\n8\n7\n6\n5\n4\n3\n2\n1\n0\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
