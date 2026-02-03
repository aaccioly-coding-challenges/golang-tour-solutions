package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestInterfacesProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "5\n"
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
		{"positive", MyFloat(3.14), 3.14},
		{"negative", MyFloat(-2.71), 2.71},
		{"zero", MyFloat(0), 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.f.Abs(); got != tc.want {
				t.Errorf("MyFloat.Abs() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestVertex_Abs(t *testing.T) {
	tests := []struct {
		name string
		v    Vertex
		want float64
	}{
		{"positive", Vertex{3, 4}, 5},
		{"negative", Vertex{-3, -4}, 5},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.v.Abs(); got != tc.want {
				t.Errorf("Vertex.Abs() = %v, want %v", got, tc.want)
			}
		})
	}
}
