package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestIPAddr_String(t *testing.T) {
	tests := []struct {
		name string
		ip   IPAddr
		want string
	}{
		{"loopback", IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{"googleDNS", IPAddr{8, 8, 8, 8}, "8.8.8.8"},
		{"empty", IPAddr{}, "0.0.0.0"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.ip.String(); got != tc.want {
				t.Errorf("IPAddr.String() = %q; want %q", got, tc.want)
			}
		})
	}
}

func TestExerciseStringerProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expectedLines := []string{
		"loopback: 127.0.0.1\n",
		"googleDNS: 8.8.8.8\n",
	}
	for _, expected := range expectedLines {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain %q, but it didn't. Got %q", expected, output)
		}
	}
}
