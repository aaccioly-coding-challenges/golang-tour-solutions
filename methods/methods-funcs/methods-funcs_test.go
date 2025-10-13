package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestMethodsProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "5\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

func TestAbs(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestAbs with 3,4", fields{3, 4}, 5},
		{"TestAbs with 5,12", fields{5, 12}, 13},
		{"TestAbs with 8,15", fields{8, 15}, 17},
		{"TestAbs with 7,24", fields{7, 24}, 25},
		{"TestAbs with 0,0", fields{0, 0}, 0},
		{"TestAbs with -3,-4", fields{-3, -4}, 5},
		{"TestAbs with -5,12", fields{-5, 12}, 13},
		{"TestAbs with 8,-15", fields{8, -15}, 17},
		{"TestAbs with -7,-24", fields{-7, -24}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vertex{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := Abs(v); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
