package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestImagesProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "(0,0)-(100,100)\n0 0 0 0\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
