package main

import (
	"runtime"
	"testing"
	"time"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
	"golang.org/x/tour/tree"
)

func TestSame(t *testing.T) {
	var tests = []struct {
		name string
		t1   *tree.Tree
		t2   *tree.Tree
		want bool
	}{
		{
			name: "Same trees (seed 1)",
			t1:   tree.New(1),
			t2:   tree.New(1),
			want: true,
		},
		{
			name: "Different trees (seed 1 vs 2)",
			t1:   tree.New(1),
			t2:   tree.New(2),
			want: false,
		},
		{
			name: "Nil trees",
			t1:   nil,
			t2:   nil,
			want: true,
		},
		{
			name: "Nil vs non-nil",
			t1:   nil,
			t2:   tree.New(1),
			want: false,
		},
		{
			name: "Non-nil vs nil",
			t1:   tree.New(1),
			t2:   nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Same(tt.t1, tt.t2); got != tt.want {
				t.Errorf("Same() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSame_NoLeak(t *testing.T) {
	initialGoroutines := runtime.NumGoroutine()

	t.Run("Functional Checks", TestSame)
	// Give goroutines a moment to exit.
	time.Sleep(50 * time.Millisecond)

	finalGoroutines := runtime.NumGoroutine()
	if finalGoroutines > initialGoroutines {
		t.Errorf("Goroutine leak detected: initial=%d, final=%d", initialGoroutines, finalGoroutines)
	}
}

func TestEquivalentBinaryTreesProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expected := "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\nSame(tree.New(1), tree.New(1)) = true\nSame(tree.New(1), tree.New(2)) = false\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
