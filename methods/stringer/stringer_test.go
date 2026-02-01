package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestStringerProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
