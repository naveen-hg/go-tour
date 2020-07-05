package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

/* Inorder travesal */
func read_tree(t *tree.Tree, ch chan int) {
	if (t != nil) {
		read_tree(t.Left, ch)
		ch <- t.Value
		read_tree(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	read_tree(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ret := true
	done := make(chan bool)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	go func() {
		for i := range ch1 {
			if i == <-ch2 {
				ret = true
			} else {
				ret = false
				break
			}
		}
		done <- true
	} ()
	<-done
	return ret
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Same tree?", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same tree?", Same(tree.New(1), tree.New(2)))
}