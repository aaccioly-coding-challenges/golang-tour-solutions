package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestInterfaceValuesProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "(&{Hello}, *main.T)\nHello\n(3.141592653589793, main.F)\n3.141592653589793\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
