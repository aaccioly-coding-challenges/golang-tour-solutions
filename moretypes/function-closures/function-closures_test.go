package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestFunctionClosures(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "0 0\n1 -2\n3 -6\n6 -12\n10 -20\n15 -30\n21 -42\n28 -56\n36 -72\n45 -90\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}

}
