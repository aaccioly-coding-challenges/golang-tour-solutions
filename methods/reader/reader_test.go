package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestReaderProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]\nb[:n] = \"Hello, R\"\nn = 6 err = <nil> b = [101 97 100 101 114 33 32 82]\nb[:n] = \"eader!\"\nn = 0 err = EOF b = [101 97 100 101 114 33 32 82]\nb[:n] = \"\"\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
