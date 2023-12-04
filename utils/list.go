package utils

import (
	"container/list"
	"fmt"
)

func NewList() {
	l := list.New()
	l.PushBack(1)
	l.PushFront(2)
	root := l.Front()
	for root != nil {
		fmt.Println(root.Value)
		root = root.Next()
	}
}
