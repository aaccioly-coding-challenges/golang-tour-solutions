package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestInterfaceValuesWithNilUnderlyingValuesProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "(<nil>, *main.T)\n<nil>\n(&{hello}, *main.T)\nhello\n"

	if output != expected {
		t.Errorf("Expected output to be '%s', but got '%s'", expected, output)
	}
}
