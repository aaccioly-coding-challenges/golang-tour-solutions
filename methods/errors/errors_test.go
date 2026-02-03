package main

import (
	"regexp"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestErrorsProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := `^at \d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d+ \+\d{4} \w+ m=\+\d+\.\d+, it didn't work\n$`

	if match, _ := regexp.MatchString(expected, output); !match {
		t.Errorf("Expected output %q to match pattern %q", output, expected)
	}

}
