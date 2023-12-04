package utils

import (
	"fmt"

	"github.com/emirpasic/gods/trees/redblacktree"
)

type pair struct {
	ID    int
	Value int
}

func RedBlackTree() {
	trees := redblacktree.NewWith(func(a, b interface{}) int {
		return a.(pair).Value - b.(pair).Value
	})
	trees.Put(pair{1, 2}, 1) // data, cnt
	trees.Put(pair{2, 3}, 1)
	node := trees.Right()
	fmt.Println(node.Key.(pair).Value)
}
