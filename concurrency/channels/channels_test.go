package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestChannelsProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expectedInOrder := "17 -5 12\n"
	expectedOutOfOrder := "-5 17 12\n"
	if output != expectedOutOfOrder && output != expectedInOrder {
		t.Errorf("Expected output %q or %q, got %q", expectedOutOfOrder, expectedInOrder, output)
	}
}
