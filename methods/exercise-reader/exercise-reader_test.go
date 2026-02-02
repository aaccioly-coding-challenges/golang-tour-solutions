package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestExerciseReaderProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "OK!\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

func TestMyReader_Read(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Buffer size 1",
			args: args{in: []byte{0}},
			want: 1,
		},
		{
			name: "Buffer size 5",
			args: args{in: []byte{0, 0, 0, 0, 0}},
			want: 5,
		},
		{
			name: "Buffer size 0",
			args: args{in: []byte{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := MyReader{}
			got, _ := r.Read(tt.args.in)
			if got != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
			for i := 0; i < got; i++ {
				if tt.args.in[i] != 'A' {
					t.Errorf("Read() buffer at index %d = %v, want %v", i, tt.args.in[i], 'A')
				}
			}
		})
	}
}
