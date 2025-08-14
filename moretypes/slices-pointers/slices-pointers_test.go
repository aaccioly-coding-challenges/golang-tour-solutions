package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSlicesPointersProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "[John Paul George Ringo]\n[John Paul] [Paul George]\n[John XXX] [XXX George]\n[John XXX George Ringo]\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
