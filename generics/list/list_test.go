package main

import (
	"reflect"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestList_Iterator(t *testing.T) {
	l := New(1)
	l.Push(2)
	l.Push(3)

	var got []int
	for v := range l.All {
		got = append(got, v)
	}

	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestList_IteratorBreak(t *testing.T) {
	l := New(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)

	var got []int
	for v := range l.All {
		got = append(got, v)
		if v == 2 {
			break
		}
	}

	want := []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestList_Size(t *testing.T) {
	l := New(1)
	if l.Size() != 1 {
		t.Errorf("expected size 1, got %d", l.Size())
	}
	l.Push(2)
	if l.Size() != 2 {
		t.Errorf("expected size 2, got %d", l.Size())
	}
	l.Push(3)
	if l.Size() != 3 {
		t.Errorf("expected size 3, got %d", l.Size())
	}
	_, _ = l.Pop()
	if l.Size() != 2 {
		t.Errorf("expected size 2, got %d", l.Size())
	}
}

func TestList_PopLast(t *testing.T) {
	l := New(1)
	l.Push(2)
	l.Push(3)

	val, err := l.Pop()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 3 {
		t.Errorf("expected 3 (last), got %v", val)
	}

	val, err = l.Pop()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 2 {
		t.Errorf("expected 2 (last), got %v", val)
	}

	val, err = l.Pop()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("expected 1 (last), got %v", val)
	}

	_, err = l.Pop()
	if err == nil {
		t.Errorf("expected error when popping from empty list")
	}
}

func TestListProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "Hello\nWorld!\nSize: 2\nPopped: World!, New Size: 1\nPopped: Hello, New Size: 0\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
