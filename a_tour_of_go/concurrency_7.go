package main

import (
	"golang.org/x/tour/tree"
)

// CheckEquivalentTrees ...
func CheckEquivalentTrees(i1, i2 int) bool {
	return doCheckEquivalentTrees(tree.New(i1), tree.New(i2))
}

func doCheckEquivalentTrees(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go walkTree(t1, ch1)
	go walkTree(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func walkTree(t *tree.Tree, ch chan int) {
	recWalkTree(t, ch)
	close(ch)
}

func recWalkTree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		recWalkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		recWalkTree(t.Right, ch)
	}
}
