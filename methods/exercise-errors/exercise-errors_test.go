package main

import (
	"math"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestExerciseErrorsProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "1.4142135623746899 <nil>\n0 cannot Sqrt negative number: -2\n"

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

func TestSqrt(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"positive square root (2)", args{2}, 1.4142135623746899, false},
		{"positive square root (4)", args{4}, 2, false},
		{"positive square root (9)", args{9}, 3, false},
		{"zero", args{0}, 0, false},
		{"negative number", args{-2}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sqrt(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sqrt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && math.Abs(got-tt.want) > 1e-10 {
				t.Errorf("Sqrt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
