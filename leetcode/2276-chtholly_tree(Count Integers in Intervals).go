package leetcode

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

type CountIntervals struct {
	*redblacktree.Tree
	cnt int
}

func Constructor() CountIntervals {
	return CountIntervals{
		redblacktree.NewWithIntComparator(),
		0,
	}
}

func (this *CountIntervals) Add(left int, right int) {
	for node, _ := this.Ceiling(left); node != nil && node.Value.(int) <= right; node, _ = this.Ceiling(left) {
		r, l := node.Key.(int), node.Value.(int)
		this.cnt -= (r - l + 1)
		if l <= left {
			left = l
		}
		if r >= right {
			right = r
		}
		this.Remove(node.Key)
	}
	this.Put(right, left)
	this.cnt += (right - left + 1)
}

func (this *CountIntervals) Count() int {
	return this.cnt
}
