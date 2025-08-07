package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSqrt(t *testing.T) {

	tests := []struct {
		name string
		x    float64
		want float64
	}{
		{"TestSqrt with 1", 1, 1.0},
		{"TestSqrt with 2", 2, 1.4142135623730951},
		{"TestSqrt with 4", 4, 2.0},
		{"TestSqrt with 9", 9, 3.0},
		{"TestSqrt with 16", 16, 4.0},
		{"TestSqrt with 25", 25, 5.0},
		{"TestSqrt with 36", 36, 6.0},
		{"TestSqrt with 49", 49, 7.0},
		{"TestSqrt with 64", 64, 8.0},
		{"TestSqrt with 81", 81, 9.0},
		{"TestSqrt with 100", 100, 10.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.x); got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewtonRaphsonProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "1.4142135623730951\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
