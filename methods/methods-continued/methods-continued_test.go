package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestMethodInNonStructTypesProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "1.4142135623730951\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
func TestMyFloat_Abs(t *testing.T) {
	tests := []struct {
		name string
		f    MyFloat
		want float64
	}{
		{"TestMyFloat_Abs with -3.14", MyFloat(-3.14), 3.14},
		{"TestMyFloat_Abs with 2.71", MyFloat(2.71), 2.71},
		{"TestMyFloat_Abs with 0", MyFloat(0), 0},
		{"TestMyFloat_Abs with -0", MyFloat(-0), 0},
		{"TestMyFloat_Abs with -123456.789", MyFloat(-123456.789), 123456.789},
		{"TestMyFloat_Abs with 123456.789", MyFloat(123456.789), 123456.789},
		{"TestMyFloat_Abs with -1e-10", MyFloat(-1e-10), 1e-10},
		{"TestMyFloat_Abs with 1e-10", MyFloat(1e-10), 1e-10},
		{"TestMyFloat_Abs with -1e10", MyFloat(-1e10), 1e10},
		{"TestMyFloat_Abs with 1e10", MyFloat(1e10), 1e10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Abs(); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
