package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestIndex(t *testing.T) {
	type args[T comparable] struct {
		s []T
		x T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[any]{
		{
			name: "int found",
			args: args[any]{s: []any{10, 20, 15, -10}, x: 15},
			want: 2,
		},
		{
			name: "int not found",
			args: args[any]{s: []any{10, 20, 15, -10}, x: 5},
			want: -1,
		},
		{
			name: "string found",
			args: args[any]{s: []any{"a", "b", "c"}, x: "b"},
			want: 1,
		},
		{
			name: "string not found",
			args: args[any]{s: []any{"a", "b", "c"}, x: "d"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Index(tt.args.s, tt.args.x); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexProgream(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "2\n-1\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
