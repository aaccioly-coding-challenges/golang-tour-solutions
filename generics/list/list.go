package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
	size int
}

func New[T any](v T) *List[T] {
	return &List[T]{val: v, size: 1}
}

func (l *List[T]) Push(v T) {
	l.size++
	curr := l
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = New[T](v)
}

func (l *List[T]) Pop() (T, error) {
	if l.size == 0 {
		return *new(T), fmt.Errorf("list is empty")
	}

	if l.next == nil {
		l.size--
		return l.val, nil
	}

	l.size--
	curr := l
	for curr.next.next != nil {
		curr = curr.next
	}

	val := curr.next.val
	curr.next = nil
	return val, nil
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) All(yield func(T) bool) {
	for curr := l; curr != nil; curr = curr.next {
		if !yield(curr.val) {
			return
		}
	}
}

func main() {
	list := New[string]("Hello")
	list.Push("World!")
	for elem := range list.All {
		fmt.Println(elem)
	}
	fmt.Printf("Size: %d\n", list.Size())
	v, _ := list.Pop()
	fmt.Printf("Popped: %s, New Size: %d\n", v, list.Size())
	v, _ = list.Pop()
	fmt.Printf("Popped: %s, New Size: %d\n", v, list.Size())
}
