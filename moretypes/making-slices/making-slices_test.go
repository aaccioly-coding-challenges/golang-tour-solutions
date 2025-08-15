package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestMakingSlicesProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "a len=5 cap=5 [0 0 0 0 0]\nb len=0 cap=5 []\nc len=2 cap=5 [0 0]\nd len=3 cap=3 [0 0 0]\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
