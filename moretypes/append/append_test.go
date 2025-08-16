package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestAppendProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "len=0 cap=0 []\nlen=1 cap=1 [0]\nlen=2 cap=2 [0 1]\nlen=5 cap=6 [0 1 2 3 4]\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
