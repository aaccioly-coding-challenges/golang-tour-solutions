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

func TestVertex_Abs(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestVertex_Abs with 3,4", fields{3, 4}, 5},
		{"TestVertex_Abs with 5,12", fields{5, 12}, 13},
		{"TestVertex_Abs with 8,15", fields{8, 15}, 17},
		{"TestVertex_Abs with 7,24", fields{7, 24}, 25},
		{"TestVertex_Abs with 0,0", fields{0, 0}, 0},
		{"TestVertex_Abs with -3,-4", fields{-3, -4}, 5},
		{"TestVertex_Abs with -5,12", fields{-5, 12}, 13},
		{"TestVertex_Abs with 8,-15", fields{8, -15}, 17},
		{"TestVertex_Abs with -7,-24", fields{-7, -24}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vertex{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Abs(); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
