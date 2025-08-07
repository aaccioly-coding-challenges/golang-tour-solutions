package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSwitchProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "Go runs on\n"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output %q to contain %q", output, expected)
	}
}
