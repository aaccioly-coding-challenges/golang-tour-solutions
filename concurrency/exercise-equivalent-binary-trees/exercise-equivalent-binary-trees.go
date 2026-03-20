package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, quit chan struct{}) {
	walkAux(t, ch, quit)
	close(ch)
}

func walkAux(t *tree.Tree, ch chan int, quit chan struct{}) {
	if t == nil {
		return
	}
	walkAux(t.Left, ch, quit)
	select {
	case ch <- t.Value:
	case <-quit:
		return
	}
	walkAux(t.Right, ch, quit)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan struct{})
	defer close(quit)

	go Walk(t1, ch1, quit)
	go Walk(t2, ch2, quit)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan struct{})
	defer close(quit)
	go Walk(tree.New(1), ch, quit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Printf("Same(tree.New(1), tree.New(1)) = %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("Same(tree.New(1), tree.New(2)) = %v\n", Same(tree.New(1), tree.New(2)))
}
