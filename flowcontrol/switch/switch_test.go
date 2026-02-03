package main

import (
	"regexp"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSwitchProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	pattern := `^Go runs on \w+\.\n$`
	matched, _ := regexp.MatchString(pattern, output)
	if !matched {
		t.Errorf("Expected output %q to match pattern %q", output, pattern)
	}
}

func Test_switchOs(t *testing.T) {
	tests := []struct {
		name string
		os   string
		want string
	}{
		{"Test_switchOs with darwin", "darwin", "macOS."},
		{"Test_switchOs with linux", "linux", "Linux."},
		{"Test_switchOs with windows", "windows", "Windows."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := switchOs(tt.os); got != tt.want {
				t.Errorf("switchOs() = %v, want %v", got, tt.want)
			}
		})
	}
}
